// Command hotreload 监听项目根目录下的 .go 文件，
// 保存后自动重新编译并重启 test.go。
//
// 用法（在 H:\project\studyGo 下执行）：
//
//	go run ./tools/hotreload
//
// 按 Ctrl+C 停止监听。
package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	buildDir = ".reload"
	debounce = 250 * time.Millisecond
)

type runner struct {
	root       string
	src        string
	bin        string
	ext        string
	lastSource []byte
	cmd        *exec.Cmd
	done       chan error
}

func newRunner() (*runner, error) {
	root, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	exe := "test"
	if runtime.GOOS == "windows" {
		exe += ".exe"
	}

	src := filepath.Join(root, "test.go")
	source, _ := os.ReadFile(src)

	return &runner{
		root:       root,
		src:        src,
		bin:        filepath.Join(root, buildDir, exe),
		ext:        filepath.Ext(exe),
		lastSource: source,
	}, nil
}

// lastDemo 读取 test.go 记录的上次选择，重启时自动重跑。
func (r *runner) lastDemo() int {
	data, err := os.ReadFile(filepath.Join(r.root, ".last_demo"))
	if err != nil {
		return 0
	}
	n, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil || n < 1 || n > 9 {
		return 0
	}
	return n
}

// build 编译项目根目录的 main 包。使用 "." 而不是写死 test.go，
// 这样以后把代码拆成多个 .go 文件也能正常工作。
func (r *runner) build() error {
	dir := filepath.Dir(r.bin)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	// 每次构建都用新文件名，避免 Windows 上旧进程还没完全释放 exe 时
	// 覆盖失败（The process cannot access the file...）。
	old, _ := filepath.Glob(filepath.Join(dir, "test-*"+r.ext))
	for _, path := range old {
		_ = os.Remove(path)
	}
	_ = os.Remove(filepath.Join(dir, "test"+r.ext))
	r.bin = filepath.Join(dir, fmt.Sprintf("test-%d%s", time.Now().UnixNano(), r.ext))

	cmd := exec.Command("go", "build", "-o", r.bin, ".")
	cmd.Dir = r.root
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (r *runner) start() error {
	var lastErr error

	var args []string
	if n := r.lastDemo(); n > 0 {
		args = []string{"-demo", strconv.Itoa(n)}
	}

	// Windows 上杀毒软件可能刚扫描完新生成的 exe，
	// 导致紧接着的启动短暂报 Access is denied，所以重试几次。
	for attempt := 0; attempt < 5; attempt++ {
		cmd := exec.Command(r.bin, args...)
		cmd.Dir = r.root
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Start(); err != nil {
			lastErr = err
			time.Sleep(250 * time.Millisecond)
			continue
		}

		r.cmd = cmd
		r.done = make(chan error, 1)
		go func(done chan<- error) {
			done <- cmd.Wait()
		}(r.done)
		return nil
	}

	return lastErr
}

func (r *runner) stop() {
	if r.cmd == nil || r.cmd.Process == nil {
		return
	}

	_ = r.cmd.Process.Kill()
	if r.done != nil {
		select {
		case <-r.done:
		case <-time.After(2 * time.Second):
			fmt.Println("警告：旧进程没有在 2 秒内退出")
		}
	}

	r.cmd = nil
	r.done = nil
}

func (r *runner) restart() {
	line := strings.Repeat("─", 60)

	fmt.Println()
	fmt.Println(paint(ansiCyan, line))
	fmt.Println(paint(ansiCyan, "[检测] test.go 发生变化"))
	fmt.Println(paint(ansiCyan, line))

	if source, err := os.ReadFile(r.src); err == nil {
		printSourceDiff(string(r.lastSource), string(source))
		r.lastSource = source
	}

	r.stop()

	fmt.Println()
	fmt.Println("[编译] 正在重新编译...")
	if err := r.build(); err != nil {
		fmt.Println(paint(ansiRed, "[编译] 失败。修复代码并保存后会自动重试。"))
		fmt.Println(paint(ansiDim, line))
		return
	}
	fmt.Println(paint(ansiGreen, "[编译] 成功"))

	if n := r.lastDemo(); n > 0 {
		fmt.Printf("[运行] 自动重跑第 %d 节（上次选择）\n", n)
	} else {
		fmt.Println("[运行] 启动菜单，等待选择")
	}
	fmt.Println(paint(ansiDim, line))

	if err := r.start(); err != nil {
		fmt.Println(paint(ansiRed, "[运行] 启动失败："), err)
	}
}

// latestChange 返回根目录下所有 .go 文件中最新的修改时间和总大小。
func (r *runner) latestChange() (time.Time, int64, error) {
	entries, err := os.ReadDir(r.root)
	if err != nil {
		return time.Time{}, 0, err
	}

	var latest time.Time
	var total int64
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			continue
		}
		if info.ModTime().After(latest) {
			latest = info.ModTime()
		}
		total += info.Size()
	}

	if latest.IsZero() {
		return time.Time{}, 0, fmt.Errorf("没有找到 .go 文件")
	}
	return latest, total, nil
}

func main() {
	initColor()

	r, err := newRunner()
	if err != nil {
		fmt.Fprintln(os.Stderr, "初始化失败：", err)
		os.Exit(1)
	}

	if _, err := os.Stat(r.src); err != nil {
		fmt.Fprintf(os.Stderr, "找不到 %s：%v\n", r.src, err)
		os.Exit(1)
	}

	fmt.Println("Go 热重载已启动")
	fmt.Println("  监听：", r.root)
	fmt.Println("  保存 .go 文件后自动重编译，并重跑上次选择的小节")
	fmt.Println("  按 Ctrl+C 停止")
	fmt.Println()

	lastMod, lastSize, err := r.latestChange()
	if err != nil {
		fmt.Fprintln(os.Stderr, "读取文件状态失败：", err)
		os.Exit(1)
	}

	fmt.Println("[启动] 首次编译...")
	if err := r.build(); err != nil {
		fmt.Println(paint(ansiRed, "[启动] 编译失败。修复代码并保存后会自动重试。"))
	} else {
		fmt.Println(paint(ansiGreen, "[启动] 编译成功"))
		if err := r.start(); err != nil {
			fmt.Println(paint(ansiRed, "[启动] 启动失败："), err)
		}
	}

	ticker := time.NewTicker(debounce)
	defer ticker.Stop()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	defer signal.Stop(sig)

	for {
		select {
		case <-ticker.C:
			mod, size, err := r.latestChange()
			if err != nil || (mod.Equal(lastMod) && size == lastSize) {
				continue
			}
			lastMod, lastSize = mod, size
			r.restart()

		case err := <-r.done:
			r.cmd = nil
			if err != nil {
				fmt.Println("\n程序退出：", err)
			} else {
				fmt.Println("\n程序已退出。")
			}
			fmt.Println("等待 .go 文件变化；按 Ctrl+C 停止监听。")

		case <-sig:
			fmt.Println("\n正在停止热重载...")
			r.stop()
			_ = os.RemoveAll(filepath.Join(r.root, buildDir))
			return
		}
	}
}
