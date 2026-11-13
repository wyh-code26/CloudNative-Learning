# 重定向与管道

## 文件描述符

| fd | 含义 |
|---|---|
| 0 | stdin 标准输入 |
| 1 | stdout 标准输出 |
| 2 | stderr 标准错误 |

## 输出重定向

```bash
cmd > file          # 覆盖 stdout
cmd >> file         # 追加 stdout
cmd 2> err.log      # stderr 到文件
cmd > out.log 2>&1  # stdout+stderr 到同一文件
cmd &> all.log      # 同上，Bash 简写
```

## 输入重定向

```bash
cmd < file
while read -r line; do echo "$line"; done < file.txt
```

## 丢弃输出

```bash
cmd > /dev/null 2>&1
```

## 管道

```bash
ls -l | grep ".md"
ps aux | grep nginx | grep -v grep
```

管道默认只传 stdout，stderr 不会进管道：

```bash
cmd 2>&1 | tee log.txt
```

## tee

```bash
cmd | tee out.log          # 同时输出到屏幕和文件
cmd | tee -a out.log       # 追加
```

## here document

```bash
cat <<EOF
hello $USER
EOF
```

禁止变量展开：

```bash
cat <<'EOF'
hello $USER
EOF
```

## here string

```bash
grep "wyh" <<< "$text"
```

## 进程替换

```bash
diff <(ls dir1) <(ls dir2)
while read -r line; do echo "$line"; done < <(ls)
```

## 常见坑

- `cmd > file 2>&1` 顺序不能反，`2>&1 > file` 只把 stdout 重定向到文件，stderr 还在屏幕。
- 管道会开子 shell，管道里 `cd`、变量赋值对外面无效。
- `>` 会立刻清空文件，不确定时先用 `>>`。