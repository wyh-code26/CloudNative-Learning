
今日重点：sed 的原地修改与正则替换

```bash
# 1. 删除特定行
sed -i '88,97d' main.go          # 删除第88到97行

# 2. 在匹配行前后插入内容
sed -i '/cred, err := vc.IssueCredential/i\    // 强制验证 ZKP 证明' main.go

# 3. 替换整个函数体（从开始到闭合括号）
sed -i '/^func tempProofPath()/,/^}/c\func tempProofPath() string {\n\treturn "new_body"\n}' main.go

# 4. 删除特定导入行
sed -i '/"encoding\/json"/d' vc/crypto.go

# 5. 确认修改后的内容
grep -A 6 "func tempProofPath" main.go  # 显示匹配行及之后6行
```

底层本质：sed 是流编辑器，逐行读入模式空间，执行编辑命令后输出。-i 直接写回原文件。/^func tempProofPath()/,/^}/c 表示从函数声明开始到闭合大括号结束，用新的文本替换整个匹配块。

对应之前所学：sed 和 grep 同属 Unix 文本处理三剑客。sed 负责改，grep 负责查。今天批量修复代码括号匹配问题，就是两者的组合——先用 grep -n 定位，再用 sed 精确修改。

⚠️ 排错预判：sed 替换后 Go 编译报错 → 大概率是括号、大括号被误伤，用 sed -n 'M,Np' 查看范围行号附近的代码，逐行数括号配对。\ 转义符过多时可写入临时 .sh 脚本执行，避免 shell 解析错误。