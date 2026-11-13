# Contributing

感谢你愿意改进这个 Go 学习项目。无论是修正错别字、补充练习、改进热重载工具，还是新增示例，都欢迎提交 Issue 或 Pull Request。

## 开发环境

- Go 1.21 或更高版本
- Git
- Windows、macOS 或 Linux

## 本地检查

提交前请运行：

```bash
gofmt -w .
go vet ./...
go test ./...
```

CI 会在 Ubuntu、Windows、macOS 上执行 `go vet`、`go test` 和 `go build`，并在 Ubuntu 上检查 `gofmt`。

## 代码约定

- Go 代码使用 `gofmt` 格式化，缩进使用 Tab。
- 示例代码以清晰为先，变量命名要能让初学者看懂。
- 新增示例请同时补充中文注释和至少一个练习。
- 修改热重载工具时，请手动验证“保存文件 → 显示 diff → 自动重跑”的流程。
- 文档使用简体中文或英文；新增文档请同步更新 README 的链接。

## 提交信息

建议使用清晰的提交信息，例如：

```text
feat: add slice deletion exercise
fix: keep last demo after hot reload
docs: improve beginner guide
```

## Pull Request

请在 PR 描述中说明：

1. 改了什么；
2. 为什么改；
3. 如何验证（命令或截图/终端输出）。

---

# Contributing (English)

Thanks for improving this Go learning project. Bug reports, documentation fixes, new exercises, and tooling improvements are all welcome.

Before opening a pull request, run:

```bash
gofmt -w .
go vet ./...
go test ./...
```

Keep examples beginner-friendly, add Chinese comments for new lessons, and update the README links when adding documentation.
