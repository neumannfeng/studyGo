package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	ansiReset = "\x1b[0m"
	ansiBold  = "\x1b[1m"
	ansiCyan  = "\x1b[36m"
	ansiRed   = "\x1b[31m"
	ansiGreen = "\x1b[32m"
	ansiDim   = "\x1b[2m"
)

var colorEnabled bool

// initColor 在 Windows 上开启控制台 ANSI 支持，并判断输出是否为终端。
func initColor() {
	if !enableANSI() {
		colorEnabled = false
		return
	}
	info, err := os.Stdout.Stat()
	colorEnabled = os.Getenv("NO_COLOR") == "" &&
		err == nil &&
		info.Mode()&os.ModeCharDevice != 0
}

func paint(code, text string) string {
	if !colorEnabled {
		return text
	}
	return code + text + ansiReset
}

type diffLine struct {
	kind  byte // ' ' 不变，'-' 删除，'+' 新增
	oldNo int
	newNo int
	text  string
}

// diffLines 用最长公共子序列比较两个文本，返回逐行差异。
func diffLines(oldText, newText string) []diffLine {
	oldLines := splitLines(oldText)
	newLines := splitLines(newText)
	n, m := len(oldLines), len(newLines)

	// dp[i][j] 表示 oldLines[i:] 与 newLines[j:] 的 LCS 长度。
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if oldLines[i] == newLines[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else if dp[i+1][j] >= dp[i][j+1] {
				dp[i][j] = dp[i+1][j]
			} else {
				dp[i][j] = dp[i][j+1]
			}
		}
	}

	lines := make([]diffLine, 0, n+m)
	i, j := 0, 0
	oldNo, newNo := 1, 1
	for i < n && j < m {
		switch {
		case oldLines[i] == newLines[j]:
			lines = append(lines, diffLine{' ', oldNo, newNo, oldLines[i]})
			i++
			j++
			oldNo++
			newNo++
		case dp[i+1][j] >= dp[i][j+1]:
			lines = append(lines, diffLine{'-', oldNo, 0, oldLines[i]})
			i++
			oldNo++
		default:
			lines = append(lines, diffLine{'+', 0, newNo, newLines[j]})
			j++
			newNo++
		}
	}
	for ; i < n; i++ {
		lines = append(lines, diffLine{'-', oldNo, 0, oldLines[i]})
		oldNo++
	}
	for ; j < m; j++ {
		lines = append(lines, diffLine{'+', 0, newNo, newLines[j]})
		newNo++
	}

	return lines
}

func splitLines(s string) []string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.TrimSuffix(s, "\n")
	if s == "" {
		return nil
	}
	return strings.Split(s, "\n")
}

func printSourceDiff(oldText, newText string) {
	lines := diffLines(oldText, newText)

	adds, dels := 0, 0
	for _, line := range lines {
		switch line.kind {
		case '+':
			adds++
		case '-':
			dels++
		}
	}
	if adds == 0 && dels == 0 {
		fmt.Println(paint(ansiDim, "[变更] 源码内容没有变化"))
		return
	}

	fmt.Printf("%s  文件：test.go    -%d 行 / +%d 行\n",
		paint(ansiBold, "[变更]"), dels, adds)
	fmt.Println(paint(ansiDim, "  行号 │ 内容"))

	// 变更行前后各保留 2 行上下文，并合并重叠区域。
	type interval struct{ start, end int }
	var spans []interval
	for i, line := range lines {
		if line.kind == ' ' {
			continue
		}

		start := i - 2
		if start < 0 {
			start = 0
		}
		end := i + 3
		if end > len(lines) {
			end = len(lines)
		}

		if len(spans) > 0 && start <= spans[len(spans)-1].end {
			if end > spans[len(spans)-1].end {
				spans[len(spans)-1].end = end
			}
			continue
		}
		spans = append(spans, interval{start, end})
	}

	for i, span := range spans {
		if i > 0 {
			fmt.Println(paint(ansiDim, "  ..."))
		}
		for _, line := range lines[span.start:span.end] {
			printDiffLine(line)
		}
	}
}

func printDiffLine(line diffLine) {
	switch line.kind {
	case '-':
		fmt.Printf("%s %4d | %s\n",
			paint(ansiRed, "-"), line.oldNo, paint(ansiRed, line.text))
	case '+':
		fmt.Printf("%s %4d | %s\n",
			paint(ansiGreen, "+"), line.newNo, paint(ansiGreen, line.text))
	default:
		fmt.Printf("%s %4d | %s\n",
			paint(ansiDim, " "), line.newNo, line.text)
	}
}
