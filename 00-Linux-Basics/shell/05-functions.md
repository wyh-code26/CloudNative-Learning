# 函数

## 定义与调用

```bash
greet() {
  echo "hello, $1"
}

greet "wyh"
```

## 两种写法

```bash
# 写法一
foo() { echo "foo"; }

# 写法二
function foo {
  echo "foo"
}
```

推荐写法一，POSIX 兼容。

## 参数

函数内部用 `$1` `$2` `$@` `$#`，和脚本参数一样。

```bash
sum() {
  echo $(($1 + $2))
}

sum 3 5      # 8
```

## 返回值

- `return` 返回的是**退出码**，0~255，不是数据。
- 要返回数据，用 `echo` 输出，调用方用命令替换接。

```bash
get_name() {
  echo "wyh"
}

name=$(get_name)
echo "$name"
```

## 局部变量

```bash
foo() {
  local x=1
  echo "$x"
}
```

不加 `local` 会污染全局。

## 检查参数

```bash
backup() {
  if [[ $# -lt 1 ]]; then
    echo "usage: backup <dir>" >&2
    return 1
  fi
  local dir="$1"
  [[ -d "$dir" ]] || { echo "not a dir: $dir" >&2; return 1; }
  echo "backup $dir"
}
```

## 常见坑

- `return` 只能返回整数退出码，不能返回字符串。
- 函数里 `local` 尽量都加，避免变量泄漏。
- 函数定义要在调用之前。