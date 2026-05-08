
今日指令：go env 与 go mod graph — Go 项目的环境诊断利器

今天在反复排查阿里云 SDK 的依赖问题时，我们一直在和 Go 的模块系统打交道。这两个命令就是用来诊断 Go 环境问题的核心工具。

```bash
# 1. go env — 查看 Go 编译器的所有环境变量
go env GOPATH      # 查看工作区路径
go env GOPROXY     # 查看模块代理地址（今天反复改的就是这个）
go env GOROOT      # 查看 Go 安装路径

# 2. go mod graph — 查看模块依赖关系图
go mod graph | grep cloudauth   # 查看 cloudauth 包的依赖链
go mod graph | head -20         # 查看依赖树的根节点

# 3. go mod why — 解释为什么需要某个包（今天没用上但很有用）
go mod why -m github.com/aliyun/alibaba-cloud-sdk-go
```

底层本质：go env 读取的是 Go 编译器的全局配置，这些配置决定了编译时去哪拉依赖、编译产物放哪、CGO 是否启用等一系列底层行为。go mod graph 则是把 go.mod 和 go.sum 里记录的依赖关系以“节点-边”的方式打印出来，本质上是一张有向图。今天阿里云 SDK 反复报 no required module 就是这张图里缺少了目标节点对应的边。

对应你的项目：今天你发现 go.mod 中 cloudauth 的记录反复丢失，就是因为 GOPROXY 指向的代理站点没有正确缓存这个包。换成 goproxy.cn 后重新拉取，依赖图里的这条边才被正确写入。

⚠️ 排错预判：go get 拉不下来包 → 先查 go env GOPROXY 确认代理地址，再查 go env GOPATH 确认工作区路径可写，最后用 go mod graph | grep 包名 确认依赖是否已写入依赖图。