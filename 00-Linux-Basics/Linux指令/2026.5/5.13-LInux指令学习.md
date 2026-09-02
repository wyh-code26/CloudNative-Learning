今日指令：VS Code 下的 Ctrl+Shift+P 与 Go 环境变量配置

今天在 Windows 上配置 VS Code 的 Go 开发环境。以下是 Go 开发必备的两条配置命令：

```bash
# 启用 Go Modules（现代 Go 项目的依赖管理方式）
go env -w GO111MODULE=on

# 设置国内代理（解决 go get 拉包慢的问题）
go env -w GOPROXY=https://goproxy.cn,direct
```

底层本质：GO111MODULE=on 强制启用 Go Modules 模式，让 Go 编译器根据 go.mod 文件管理依赖，而非旧的 GOPATH 模式。GOPROXY 指定模块代理服务器，direct 表示代理不可用时回源拉取。这两条配置决定了你的 go mod tidy 和 go build 能否顺利执行。

对应之前所学：这和 Linux 下的 export PATH=$PATH:/usr/local/go/bin 同类——都是环境变量配置，只不过 Go 的环境变量通过 go env -w 持久化，而 Linux 的 export 只在当前 Shell 会话生效。

⚠️ 排错预判：VS Code 报 gopls 未安装 → 打开任意 .go 文件，点右下角弹出的 "Install All" 安装全部 Go 工具；go mod tidy 拉包超时 → 先执行 go env -w GOPROXY=https://goproxy.cn,direct 切换代理。