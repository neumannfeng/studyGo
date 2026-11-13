# Beginner Guide: Run the Go Learning Lab from Scratch

[中文](beginner-guide.zh-CN.md) · [English](beginner-guide.en.md) · [Back to README](../README.en.md) · [Learning Plan (Chinese)](../学习计划.md)

> Goal: install Go, run the project, edit code, and publish it to GitHub in 30–60 minutes.
> Works on Windows, macOS, and Linux. Platform-specific commands are marked.

---

## Contents

1. [What you will accomplish](#1-what-you-will-accomplish)
2. [Install Go](#2-install-go)
3. [Get the project](#3-get-the-project)
4. [Run it for the first time](#4-run-it-for-the-first-time)
5. [Edit code and see the result](#5-edit-code-and-see-the-result)
6. [Use the learning menu](#6-use-the-learning-menu)
7. [Common errors](#7-common-errors)
8. [Publish to GitHub](#8-publish-to-github)
9. [Next steps](#9-next-steps)
10. [Quick command card](#10-quick-command-card)

---

## 1. What you will accomplish

By the end of this guide you will have:

- a runnable Go learning project;
- a save-to-run development loop;
- a working understanding of the menu and exercises in `test.go`;
- your own GitHub repository.

You do not need prior Go experience. You should be comfortable opening a terminal, copying commands, and editing a text file.

---

## 2. Install Go

### Windows

1. Open <https://go.dev/dl/>.
2. Download the Windows `.msi` installer.
3. Run it and keep the default options.
4. Close every terminal, then open PowerShell again.
5. Check:

```powershell
go version
```

You should see something like `go version go1.27.3 windows/amd64`.

### macOS

Option 1: download the official `.pkg` from <https://go.dev/dl/>.

Option 2: use Homebrew:

```bash
brew install go
```

Check:

```bash
go version
```

### Linux

Follow the official instructions: <https://go.dev/doc/install>

After extracting the archive to `/usr/local`, add `/usr/local/go/bin` to `PATH`. Check:

```bash
go version
```

### If `go: command not found`

- Close and reopen the terminal;
- log out and back in if needed;
- on Windows, confirm the Go `bin` directory is in `Path`;
- on macOS/Linux, confirm `/usr/local/go/bin` is in `PATH`.

---

## 3. Get the project

### Option A: clone with Git (recommended)

```bash
git clone https://github.com/neumannfeng/studyGo.git
cd studyGo
```

The URL above uses the username `neumannfeng` and the repository name `studyGo`. If you rename the repository, replace `studyGo` accordingly. If the project is not on GitHub yet, use option B.

### Option B: download a ZIP

1. Open the GitHub repository page;
2. click **Code → Download ZIP**;
3. extract the archive;
4. open a terminal in the extracted directory:

```powershell
cd path\to\studyGo
```

### Check the directory

```bash
ls          # macOS / Linux
dir         # Windows PowerShell / CMD
```

You should see `test.go`, `go.mod`, `README.md`, and `学习计划.md`.

---

## 4. Run it for the first time

From the project root:

```bash
go run test.go
```

The first build may take a few seconds. You should see:

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

Suggested first run:

1. enter `9` to run every example;
2. enter `1` to study variables carefully;
3. enter `q` to quit.

If there is a compile error, the terminal shows the file, line, and reason. Read the first error first; it usually points to the real problem.

---

## 5. Edit code and see the result

### 5.1 Open the file

Open [test.go](../test.go) in VS Code, GoLand, or any text editor.

Find this line in lesson 1 (the line number may change):

```go
age := 8
```

Change it to:

```go
age := 9
```

Save the file.

### 5.2 Option 1: run manually

Return to the terminal, press `Ctrl+C` or enter `q` to stop the current program, then run:

```bash
go run test.go
```

Enter `1` and you will see `age = 9`.

This works, but you have to quit, rebuild, and choose the lesson every time.

### 5.3 Option 2: hot reload (recommended)

Start the watcher.

Windows:

```powershell
.\dev.bat
```

macOS / Linux:

```bash
go run ./tools/hotreload
```

After `Go 热重载已启动` appears, select lesson `1` once so the tool remembers your choice.

Now change `age := 9` to `10` and save. The terminal shows:

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

The lesson reruns automatically; you do not need to type `1` again.

### 5.4 Hot-reload keys

| Action | Effect |
| --- | --- |
| Save a `.go` file | Rebuild and rerun the last lesson |
| `q` | Quit the current program; the watcher keeps running |
| `Ctrl+C` | Stop the watcher |
| Delete `.last_demo` | Disable automatic reruns on the next start |

### 5.5 Why not true hot patching?

Go is compiled. The running program is a binary and cannot receive new source code. The watcher:

1. detects the file change;
2. prints a source diff;
3. runs `go build` again;
4. starts a new process and passes `-demo N` to rerun the last lesson.

In-memory state is reset. This is expected.

### 5.6 What if compilation fails?

If you see:

```text
[编译] 失败。修复代码并保存后会自动重试。
```

fix the syntax or type error and save again. You do not need to restart the watcher. See the next section for common errors.

---

## 6. Use the learning menu

| # | Topic | Suggested timing |
| --- | --- | --- |
| 1 | Variables, constants, basic types | Week 1, Day 1–2 |
| 2 | Conditions, loops, switch | Week 1, Day 3 |
| 3 | Functions, multiple returns, closures, defer | Week 1, Day 4 |
| 4 | Slices, maps, range | Week 1, Day 5–6 |
| 5 | Structs, methods, pointers | Week 2, Day 1–2 |
| 6 | Interfaces, errors, errors.Is/As | Week 2, Day 3–4 |
| 7 | Generics | Late week 2 or week 3 |
| 8 | goroutine, channel, WaitGroup | Week 4 |
| 9 | Run every example | Quick overview |

Every lesson ends with exercises. Try them first, then compare with the code. If you are stuck for more than 20 minutes, look up the topic or ask a question.

The full schedule is in the [4-week learning plan (Chinese)](../学习计划.md).

---

## 7. Common errors

### `go: command not found` / `'go' is not recognized`

Go is not installed correctly, or `PATH` has not refreshed. Reinstall and reopen the terminal.

### `declared and not used`

You declared a variable but never used it. Go treats this as a compile error, not a warning. Remove it or use it.

### `imported and not used`

You imported a package but did not use it. Remove the import or use the package.

### `undefined: xxx`

A typo, a scope problem, or a function that has not been defined yet.

### `expected ';'` / `unexpected newline`

Usually an unclosed parenthesis, brace, or quote. Check the line before the reported error.

### `Access is denied` / cannot delete `.reload`

An old test process or antivirus scanner still holds the binary. Press `Ctrl+C` to stop the watcher, close the terminal, and try again. Each build uses a new filename, so a locked old binary does not block the next build.

### `dev.bat` does not run

In PowerShell, use `.\dev.bat`, not `dev.bat`. If batch files are restricted, run:

```bash
go run ./tools/hotreload
```

### Nothing happens after saving

- Confirm `dev.bat` or `go run ./tools/hotreload` is still running;
- edit a `.go` file in the project root;
- `q` only quits the program; it does not stop the watcher;
- `Ctrl+C` stops the watcher.

### No colors

Colors are enabled only for terminals that support ANSI. The watcher tries to enable ANSI on classic Windows consoles and falls back to plain `-` / `+` markers. Setting `NO_COLOR` disables colors.

---

## 8. Publish to GitHub

The examples below use the repository name `studyGo` and branch `main`.

### 8.1 Create the repository

1. Open <https://github.com/> and sign in;
2. click **+ → New repository**;
3. set Repository name to `studyGo`;
4. choose Public or Private;
5. do **not** add a README, `.gitignore`, or license, because the local project already has them;
6. click **Create repository**.

### 8.2 Configure Git

For first-time Git users:

```bash
git config --global user.name "Your Name"
git config --global user.email "you@example.com"
```

Check:

```bash
git config --global user.name
git config --global user.email
```

### 8.3 Initialize and commit

From the project root:

```bash
git init -b main
git status
git add .
git status
git commit -m "feat: initial Go learning lab"
```

If your Git version does not support `git init -b main`:

```bash
git init
git branch -M main
```

Before committing, inspect `git status` and make sure `.reload/`, `.last_demo`, and `.idea/` are not included.

### 8.4 Connect the remote

Copy the repository URL from GitHub, for example:

```text
https://github.com/neumannfeng/studyGo.git
```

Then:

```bash
git remote add origin https://github.com/neumannfeng/studyGo.git
git remote -v
git push -u origin main
```

If `origin` already exists:

```bash
git remote set-url origin https://github.com/neumannfeng/studyGo.git
```

### 8.5 GitHub authentication

GitHub does not accept your account password for Git pushes. Use one of these:

- **HTTPS + Personal Access Token**: create a token under `Settings → Developer settings → Personal access tokens`, then paste it when Git asks for a password;
- **SSH**: add your public key to GitHub and use `git@github.com:neumannfeng/studyGo.git`.

On Windows, Git Credential Manager may open a browser login window during the first push.

### 8.6 If the remote already has a README

If you created the repository with a README, the remote has an initial commit. Pull and rebase first:

```bash
git pull --rebase origin main
git push -u origin main
```

Resolve conflicts if prompted, then continue with `git add` and `git rebase --continue`.

### 8.7 Check the result

Refresh the GitHub repository page and confirm:

- the README renders correctly;
- `test.go`, `学习计划.md`, and `docs/` are present;
- `.reload/`, `.last_demo`, and `.idea/` are absent;
- the Actions tab starts running CI (the first run may take a few minutes).

### 8.8 Future updates

```bash
git add .
git commit -m "docs: update notes"
git push
```

### 8.9 GitHub CLI (optional)

If you install [GitHub CLI](https://cli.github.com/) and run `gh auth login`, you can create and push in one command:

```bash
gh repo create studyGo --public --source=. --remote=origin --push
```

---

## 9. Next steps

1. Follow the [4-week learning plan (Chinese)](../学习计划.md);
2. modify the examples and finish the exercises in each lesson;
3. build each weekly project under `practice/` and give it a short README;
4. commit regularly with Git;
5. after four weeks, turn the repository into a portfolio project.

---

## 10. Quick command card

```bash
# Interactive menu
go run test.go

# Run lesson 1 directly
go run test.go -demo 1

# Hot reload
.\dev.bat                    # Windows
go run ./tools/hotreload     # macOS / Linux

# Format, check, test
gofmt -w test.go
go vet ./...
go test ./...

# First GitHub push
git init -b main
git add .
git commit -m "feat: initial Go learning lab"
git remote add origin https://github.com/neumannfeng/studyGo.git
git push -u origin main

# Future updates
git add .
git commit -m "feat: add week 1 project"
git push
```

When something fails, read the first error in the terminal, then check [Common errors](#7-common-errors). If you open an issue, include your OS, Go version, the exact command, and the full error output.
