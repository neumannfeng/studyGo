# practice

这里放四个周项目。每个项目使用独立子目录和独立 `package main`，避免和根目录的 [test.go](../test.go) 冲突。

建议结构：

```text
practice/
├── week1-temperature/
│   └── main.go
├── week2-contacts/
│   └── main.go
├── week3-todo/
│   ├── main.go
│   └── todos.json      # 运行时生成，已加入 .gitignore
└── week4-healthcheck/
    └── main.go
```

运行示例：

```bash
go run ./practice/week1-temperature
go run ./practice/week2-contacts
go run ./practice/week3-todo
go run ./practice/week4-healthcheck
```

项目要求见 [学习计划](../学习计划.md)。第一次做项目时，先把功能跑通，再补错误处理、测试和 README。
