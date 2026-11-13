# 新手指引：从零运行 Go 学习实验室

[中文](beginner-guide.zh-CN.md) · [English](beginner-guide.en.md) · [返回 README](../README.md) · [学习计划](../学习计划.md)

> 目标：在 30–60 分钟内完成安装、运行、修改代码，并知道如何上传到 GitHub。
> 适用：Windows、macOS、Linux。命令会分别标注平台差异。

---

## 目录

1. [你会完成什么](#1-你会完成什么)
2. [安装 Go](#2-安装-go)
3. [获取项目](#3-获取项目)
4. [第一次运行](#4-第一次运行)
5. [修改代码，观察变化](#5-修改代码观察变化)
6. [使用学习菜单](#6-使用学习菜单)
7. [常见错误](#7-常见错误)
8. [上传到 GitHub](#8-上传到-github)
9. [下一步](#9-下一步)
10. [快速命令卡](#10-快速命令卡)

---

## 1. 你会完成什么

完成本指引后，你会得到：

- 一个能运行的 Go 学习项目；
- 一个可以保存后自动重跑的开发环境；
- 对 `test.go` 菜单和练习的使用方法；
- 一个属于自己的 GitHub 仓库。

你不需要提前会 Go，但需要会打开终端、复制命令、编辑文本文件。

---

## 2. 安装 Go

### Windows

1. 打开 <https://go.dev/dl/>。
2. 下载 Windows 的 `.msi` 安装包。
3. 双击安装，保持默认选项。
4. 安装完成后，关闭所有终端，再重新打开 PowerShell。
5. 运行：

```powershell
go version
```

看到类似 `go version go1.27.3 windows/amd64` 就成功了。

### macOS

方式一：下载官方 `.pkg` 安装包：<https://go.dev/dl/>

方式二：使用 Homebrew：

```bash
brew install go
```

验证：

```bash
go version
```

### Linux

推荐按官方文档安装：<https://go.dev/doc/install>

下载 tar.gz 后解压到 `/usr/local`，然后把 `/usr/local/go/bin` 加入 `PATH`。验证：

```bash
go version
```

### 如果提示 `go: command not found`

- 关闭并重新打开终端；
- 确认安装完成后重新登录；
- Windows 检查“系统属性 → 环境变量 → Path”是否包含 Go 的 `bin` 目录；
- macOS/Linux 检查 `PATH` 是否包含 `/usr/local/go/bin`。

---

## 3. 获取项目

### 方式 A：使用 Git 克隆（推荐）

```bash
git clone https://github.com/neumannfeng/studyGo.git
cd studyGo
```

上面的地址使用用户名 `neumannfeng` 和仓库名 `studyGo`。如果你改了仓库名，请把 `studyGo` 换成实际名称；如果项目还没上传到 GitHub，可以先用方式 B。

### 方式 B：下载 ZIP

1. 打开 GitHub 仓库页面；
2. 点击 **Code → Download ZIP**；
3. 解压；
4. 在终端进入解压后的目录：

```powershell
cd path\to\studyGo
```

### 检查目录

```bash
ls          # macOS / Linux
dir         # Windows PowerShell / CMD
```

应该能看到 `test.go`、`go.mod`、`README.md`、`学习计划.md`。

---

## 4. 第一次运行

在项目根目录执行：

```bash
go run test.go
```

第一次运行会先编译，可能需要几秒钟。成功后会看到：

```text
════════════════════════════════════════════════════════════
  Go 语言学习实验室 · test.go
════════════════════════════════════════════════════════════

──────────────────────── 选择学习内容 ────────────────────────
  1. 变量、常量与基本类型
  2. 条件、循环与 switch
  ...
  9. 运行全部示例
  q. 退出
请选择>
```

建议：

1. 输入 `9`，把全部示例跑一遍；
2. 输入 `1`，从变量开始仔细看；
3. 输入 `q` 退出。

如果代码有编译错误，终端会显示错误行号和原因。先读第一条错误，通常就能定位问题。

---

## 5. 修改代码，观察变化

### 5.1 打开文件

用 VS Code、GoLand 或任意文本编辑器打开 [test.go](../test.go)。

找到第 1 节里的这一行（行号可能因修改而变化）：

```go
age := 8
```

把它改成：

```go
age := 9
```

保存文件。

### 5.2 方式一：手动重新运行

回到终端，按 `Ctrl+C` 结束当前程序，或者输入 `q` 退出，然后再次运行：

```bash
go run test.go
```

输入 `1`，你会看到 `age = 9`。

这种方式最直接，但每次都要退出、重新编译、重新选择。

### 5.3 方式二：热重载（推荐）

先启动热重载。

Windows：

```powershell
.\dev.bat
```

macOS / Linux：

```bash
go run ./tools/hotreload
```

看到 `Go 热重载已启动` 后，先手动输入一次 `1`，让工具记住你的选择。

然后回到编辑器，把 `age := 9` 改成 `10` 并保存。终端会自动显示：

```text
────────────────────────────────────────────────────────────
[检测] test.go 发生变化
────────────────────────────────────────────────────────────
[变更]  文件：test.go    -1 行 / +1 行
  行号 │ 内容
   167 |
   168 | 	var name string = "Gopher"
-  169 | 	age := 9
+  169 | 	age := 10
   170 | 	height := 1.75
   171 | 	isReady := true

[编译] 成功
[运行] 自动重跑第 1 节（上次选择）
```

随后会自动重跑第 1 节，不需要再输入 `1`。

### 5.4 热重载的按键

| 操作 | 效果 |
| --- | --- |
| 保存 `.go` 文件 | 自动重编译并重跑上次小节 |
| `q` | 只退出当前程序，监听继续 |
| `Ctrl+C` | 停止热重载 |
| 删除 `.last_demo` | 下次启动不自动重跑 |

### 5.5 为什么不能直接“热更新”？

Go 是编译型语言，运行的是编译后的二进制文件。源码改了，已经运行的进程不会变。热重载工具的做法是：

1. 检测文件变化；
2. 显示源码 diff；
3. 重新 `go build`；
4. 启动新进程，并传入 `-demo N` 重跑上次小节。

所以内存里的变量会重置，这是正常的。

### 5.6 编译错误怎么办？

如果保存后看到：

```text
[编译] 失败。修复代码并保存后会自动重试。
```

说明代码有语法或类型错误。修复后直接保存，不需要重启热重载工具。

常见错误见下一节。

---

## 6. 使用学习菜单

| 编号 | 主题 | 建议时间 |
| --- | --- | --- |
| 1 | 变量、常量与基本类型 | 第 1 周 Day 1–2 |
| 2 | 条件、循环与 switch | 第 1 周 Day 3 |
| 3 | 函数、多返回值、闭包与 defer | 第 1 周 Day 4 |
| 4 | 切片、map 与 range | 第 1 周 Day 5–6 |
| 5 | 结构体、方法与指针 | 第 2 周 Day 1–2 |
| 6 | 接口、错误处理与 errors.Is/As | 第 2 周 Day 3–4 |
| 7 | 泛型 | 第 2 周后段或第 3 周 |
| 8 | goroutine、channel 与 WaitGroup | 第 4 周 |
| 9 | 运行全部示例 | 快速总览 |

每节末尾都有练习。先自己做，再对照代码；卡住超过 20 分钟，再查资料或提问。

完整安排见 [学习计划](../学习计划.md)。

---

## 7. 常见错误

### `go: command not found` / `go 不是内部或外部命令`

Go 没安装好，或 `PATH` 没生效。重新安装并重开终端。

### `declared and not used`

声明了变量但没用。Go 会把它当成编译错误，不是警告。删掉变量或使用它。

### `imported and not used`

导入了包但没使用。删掉 import，或补上使用代码。

### `undefined: xxx`

名字拼错、作用域不对，或函数还没定义。

### `expected ';'` / `unexpected newline`

通常是括号、花括号或引号没闭合。看错误提示的前一行。

### `Access is denied` / `.reload` 删除失败

旧的测试进程或杀毒软件还占用着二进制文件。按 `Ctrl+C` 停止 watcher，关闭相关终端后再删除 `.reload`。热重载每次构建都会使用新文件名，不会因此无法编译。

### `dev.bat` 无法执行

PowerShell 中要写 `.\dev.bat`，不能只写 `dev.bat`。如果仍然失败，直接运行：

```bash
go run ./tools/hotreload
```

### 保存后没有自动运行

- 确认 `dev.bat` 或 `go run ./tools/hotreload` 仍在运行；
- 确认修改的是项目根目录下的 `.go` 文件；
- 输入 `q` 只会退出程序，不会停止监听；
- 按 `Ctrl+C` 才是停止监听。

### 没有颜色

颜色只在支持 ANSI 的终端启用。Windows 传统控制台会尝试自动开启；失败时会保留 `-` / `+` 标记。设置 `NO_COLOR` 环境变量也会关闭颜色。

---

## 8. 上传到 GitHub

下面以仓库名 `studyGo`、分支 `main` 为例。

### 8.1 注册并创建仓库

1. 打开 <https://github.com/> 并登录；
2. 点击右上角 **+ → New repository**；
3. Repository name 填 `studyGo`；
4. 选择 Public 或 Private；
5. **不要**勾选 “Add a README file”“Add .gitignore”“Choose a license”，因为本地已经有这些文件；
6. 点击 **Create repository**。

### 8.2 配置 Git 身份

第一次使用 Git 时设置：

```bash
git config --global user.name "你的名字"
git config --global user.email "你的邮箱"
```

检查：

```bash
git config --global user.name
git config --global user.email
```

### 8.3 初始化并提交

在项目根目录执行：

```bash
git init -b main
git status
git add .
git status
git commit -m "feat: initial Go learning lab"
```

如果 Git 版本较旧，不支持 `git init -b main`：

```bash
git init
git branch -M main
```

提交前检查 `git status`，确保没有把 `.reload/`、`.last_demo`、`.idea/` 加进去。

### 8.4 连接远程仓库

在 GitHub 创建仓库后，复制仓库地址，例如：

```text
https://github.com/neumannfeng/studyGo.git
```

执行：

```bash
git remote add origin https://github.com/neumannfeng/studyGo.git
git remote -v
git push -u origin main
```

如果提示 `remote origin already exists`：

```bash
git remote set-url origin https://github.com/neumannfeng/studyGo.git
```

### 8.5 GitHub 认证

GitHub 不接受账户密码直接推送。推荐两种方式：

- **HTTPS + Personal Access Token**：在 GitHub `Settings → Developer settings → Personal access tokens` 创建 token，推送时密码位置填 token；
- **SSH**：在 GitHub 添加本机 SSH 公钥，远程地址使用 `git@github.com:neumannfeng/studyGo.git`。

第一次推送时，Windows 可能弹出 Git Credential Manager 登录窗口，按提示登录即可。

### 8.6 如果远程仓库已经有 README

如果你创建仓库时勾选了 README，远程会有一个初始提交。先拉取并变基：

```bash
git pull --rebase origin main
git push -u origin main
```

如果出现冲突，先解决冲突，再 `git add`、`git rebase --continue`。

### 8.7 检查上传结果

刷新 GitHub 仓库页面，确认：

- README 能正常显示；
- `test.go`、`学习计划.md`、`docs/` 都在；
- 没有 `.reload/`、`.last_demo`、`.idea/`；
- Actions 页面里的 CI 开始运行（首次可能需要几分钟）。

### 8.8 后续更新

```bash
git add .
git commit -m "docs: update notes"
git push
```

### 8.9 使用 GitHub CLI（可选）

如果安装了 [GitHub CLI](https://cli.github.com/) 并完成 `gh auth login`，可以一条命令创建并推送：

```bash
gh repo create studyGo --public --source=. --remote=origin --push
```

---

## 9. 下一步

1. 按 [学习计划](../学习计划.md) 开始第 1 周；
2. 每完成一节课，修改 `test.go` 里的示例并做练习；
3. 周末把项目放到 `practice/`，写一个简短 README；
4. 学会用 `git commit` 记录每次进步；
5. 四周后把仓库整理成作品集。

---

## 10. 快速命令卡

```bash
# 运行菜单
go run test.go

# 直接运行第 1 节
go run test.go -demo 1

# 热重载
.\dev.bat                    # Windows
go run ./tools/hotreload     # macOS / Linux

# 格式化、检查、测试
gofmt -w test.go
go vet ./...
go test ./...

# Git 首次上传
git init -b main
git add .
git commit -m "feat: initial Go learning lab"
git remote add origin https://github.com/neumannfeng/studyGo.git
git push -u origin main

# 后续更新
git add .
git commit -m "feat: add week 1 project"
git push
```

遇到问题时，先看终端的第一条错误，再看本文的[常见错误](#7-常见错误)，最后到 GitHub Issues 提问。提问时附上：操作系统、Go 版本、完整命令、完整错误信息。
