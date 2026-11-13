# 循环

## for 遍历列表

```bash
for i in 1 2 3; do
  echo "$i"
done
```

## for 遍历文件（glob）

```bash
for f in *.txt; do
  [[ -e "$f" ]] || continue      # 没有匹配时跳过
  echo "$f"
done
```

## C 风格 for

```bash
for ((i=0; i<5; i++)); do
  echo "$i"
done
```

## for 遍历命令输出

```bash
while IFS= read -r line; do
  echo "$line"
done < <(ls -1)
```

## while

```bash
n=0
while [[ $n -lt 3 ]]; do
  echo "$n"
  ((n++))
done
```

## until

```bash
n=0
until [[ $n -ge 3 ]]; do
  echo "$n"
  ((n++))
done
```

## 读取文件

```bash
while IFS= read -r line; do
  echo "$line"
done < file.txt
```

## break / continue

```bash
for i in {1..10}; do
  [[ $i -eq 3 ]] && continue     # 跳过 3
  [[ $i -eq 7 ]] && break        # 7 时退出
  echo "$i"
done
```

## 遍历数组

```bash
arr=(a "b c" d)
for x in "${arr[@]}"; do
  echo "[$x]"
done
```

## 常见坑

- `for f in $(ls)` 遇到空格文件名会出错，用 `for f in *` 或 `while read`。
- `while read` 不写 `IFS=` 会丢掉行首行尾空白。
- 不写 `-r` 会把反斜杠当转义。
- 循环里修改外部变量，管道会开子 shell，改用 `< <(...)` 或 `while ... done < file`。