# 引号与转义

## 三种引号

| 写法 | 行为 |
|---|---|
| `'...'` | 单引号：全部原样，不展开变量、不执行命令 |
| `"..."` | 双引号：展开变量、命令替换、算术，但保留空格 |
| `` `...` `` | 反引号：命令替换，等价于 `$(...)`，不推荐 |
| `$'...'` | ANSI-C 引号，支持 `\n` `\t` 等转义 |

## 示例

```bash
name="wyh"
echo '$name'        # 输出 $name
echo "$name"        # 输出 wyh
echo "a b"          # 一个参数：a b
echo a b            # 两个参数：a b
```

## 命令替换

```bash
now=$(date +%F)     # 推荐
now=`date +%F`      # 旧写法，不推荐
echo "$now"
```

## 转义

```bash
echo "He said \"hi\""    # 双引号内转义双引号
echo 'It'\''s ok'        # 单引号内嵌单引号
echo -e "a\tb\n"         # -e 让 echo 解释 \t \n
printf "a\tb\n"          # printf 更可移植
```

## 命令替换的坑

```bash
files=$(ls)             # 不推荐，文件名含空格/换行会出错
find . -type f -print0 | xargs -0 ls   # 更稳
```

## 数组与引号

```bash
arr=(a "b c" d)
for x in "${arr[@]}"; do echo "[$x]"; done   # 正确：保留每个元素
for x in ${arr[@]}; do echo "[$x]"; done     # 错误：b c 被拆开
```

## 常见坑

- `"$@"` 保留参数边界，`"$*"` 会拼成一个字符串。
- 变量引用不加双引号，遇到空格会被拆成多个参数。
- 文件名含空格时，`for f in $(ls)` 会出错，优先用 glob：`for f in *.txt`。