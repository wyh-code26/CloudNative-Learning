# Shell 基础与脚本

本目录记录 Bash/Shell 语法与脚本练习，按主题组织。

## 目录

| 文件 | 内容 |
|---|---|
| [01-variables.md](01-variables.md) | 变量、环境变量、特殊变量 |
| [02-quoting.md](02-quoting.md) | 单引号、双引号、命令替换、转义 |
| [03-conditionals.md](03-conditionals.md) | if、test、[ ]、[[ ]]、case |
| [04-loops.md](04-loops.md) | for、while、until、break、continue |
| [05-functions.md](05-functions.md) | 函数定义、参数、返回值、局部变量 |
| [06-redirection-pipes.md](06-redirection-pipes.md) | 重定向、管道、tee、here document |
| [scripts/](scripts/) | 实际脚本练习 |
| [07-shortcuts.md](07-shortcuts.md) | Bash 命令行编辑快捷键 |

## 约定

- 脚本首行统一使用 `#!/usr/bin/env bash`
- 脚本开头建议加 `set -euo pipefail`
- 变量引用统一加双引号：`"$var"`
- 练习脚本放在 `scripts/` 下