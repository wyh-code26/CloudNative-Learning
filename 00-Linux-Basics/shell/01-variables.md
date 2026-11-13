# 变量

## 定义与使用

```bash
name="wyh"          # 等号两边不能有空格
echo "$name"        # 使用变量加 $
echo "${name}"      # 推荐写法，边界清晰
```

## 只读变量

```bash
readonly PI=3.14
PI=3.15             # 报错：readonly variable
```

## 删除变量

```bash
unset name
```

## 环境变量

```bash
export NAME="wyh"           # 导出为环境变量，子进程可见
env                         # 查看所有环境变量
printenv PATH               # 查看单个
```

临时对一条命令生效：

```bash
NAME="tmp" ./script.sh
```

## 特殊变量

| 变量 | 含义 |
|---|---|
| `$0` | 脚本名 |
| `$1` ~ `$9` | 第 1~9 个位置参数 |
| `${10}` | 第 10 个参数，超过 9 必须加花括号 |
| `$#` | 参数个数 |
| `$@` | 所有参数，作为独立单词 |
| `$*` | 所有参数，作为一个整体 |
| `$?` | 上一条命令的退出码 |
| `$$` | 当前 shell 的 PID |
| `$!` | 最近一个后台进程的 PID |

## 默认值与判空

```bash
echo "${name:-default}"     # name 为空时用 default
echo "${name:=default}"     # name 为空时赋值并返回
echo "${name:?未设置}"       # name 为空时报错退出
echo "${name:+yes}"         # name 非空时返回 yes
```

## 字符串操作

```bash
str="hello world"
echo "${#str}"              # 长度：11
echo "${str:0:5}"           # 截取：hello
echo "${str/world/bash}"    # 替换第一个：hello bash
echo "${str//l/L}"          # 全部替换：heLLo worLd
echo "${str#hello }"        # 从头删最短匹配：world
echo "${str% world}"        # 从尾删最短匹配：hello
```

## 算术运算

```bash
a=1; b=2
echo $((a + b))             # 3
((a++))                     # a 自增
let "c = a * b"
```

## 常见坑

- `name = "wyh"` 会报错，等号两边不能有空格。
- `$var` 后面紧跟字母时要用 `${var}`。
- 未定义变量在 `set -u` 下会报错，建议用 `${var:-}` 兜底。