今日指令：go mod — Go 项目的模块管理核心

 mini-apiserver 项目是怎么编译的。它背后依赖的就是 go mod。它用来管理项目的模块依赖和包导入路径。

```bash
# 1. 初始化一个新模块（生成 go.mod 文件）
go mod init github.com/wuyuhang/midun  # 模块名建议用仓库路径

# 2. 自动添加缺失的依赖，移除无用的依赖
go mod tidy

# 3. 查看当前项目的所有依赖
go list -m all

# 4. 下载依赖到本地缓存
go mod download
```

底层本质：go.mod 文件定义了模块路径（就像你项目的唯一ID）和依赖版本。你在代码里写的 import "github.com/wuyuhang/midun/vc"，Go 编译器就是通过 go.mod 里的模块名来找到对应的包的。

对应的项目：《密盾》的所有内部包都整齐地放在 vc/、audit/、zkp/ 文件夹下，go.mod 文件确保了它们能互相引用，也方便别人拉取你的代码。