今日命令：golangci-lint — Go 代码的“全家桶体检”

```bash
# 1. 安装 golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# 2. 在项目根目录运行，自动检查代码风格、潜在 bug 和安全问题
cd ~/midun
golangci-lint run ./...

# 3. 只运行特定的检查器，比如只检查安全问题（gosec）
golangci-lint run --disable-all --enable gosec ./...

# 4. 查看所有可用的检查器
golangci-lint linters
```

底层本质：golangci-lint 是一个聚合工具，内部集成了多个静态分析器（linter）。每个 linter 都是一套规则引擎，逐行扫描代码的抽象语法树，匹配已知的“错误模式”或“不规范的写法”。它不运行你的代码，只读代码文本，所以叫“静态分析”。

对应之前所学：和 grep 一样，linter 也是基于模式匹配——只不过 grep 匹配的是文本，linter 匹配的是代码结构。你昨天用 grep -n 定位代码行，今天用 linter 自动发现代码问题，两者本质都是搜索。

排错预判：golangci-lint 报 command not found → Go 的 bin 目录不在 PATH 里，执行 export PATH=$PATH:$(go env GOPATH)/bin 后重试。