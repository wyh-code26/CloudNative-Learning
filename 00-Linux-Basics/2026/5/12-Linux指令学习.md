今日指令：certutil 与 powershell -Command — Windows 下的 Base64 编码方案

昨天在 Windows 下测试快瞳 OCR 接口，需要将身份证图片转成 Base64。Windows 下没有 base64 命令，但有两条替代路径：

```bash
# 方案一：certutil（CMD 自带，但输出包含头尾标记）
certutil -encode 原始文件.jpg 输出文件.txt
# 生成的 txt 文件头部有 -----BEGIN CERTIFICATE-----
# 尾部有 -----END CERTIFICATE-----
# 用于 API 请求时需要手动删除这两行和所有换行符

# 方案二：PowerShell 内联 Base64（无多余标记）
powershell -Command "[Convert]::ToBase64String([IO.File]::ReadAllBytes('C:\Users\17808\Desktop\test.jpg'))"
# 输出为纯 Base64 字符串，不包含任何头尾标记
# 适合直接拼接到 JSON 请求体中
```

底层本质：Linux 下的 base64 命令是 GNU coreutils 的一部分，直接对标准输入/输出流进行编码。Windows 没有等价的原生命令，certutil 原本是证书管理工具，-encode 只是它的副业。PowerShell 的方案更接近编程语言的库调用，[Convert]::ToBase64String() 直接调用 .NET Framework 的 Base64 编码器。

⚠️ 排错预判：certutil 生成的 Base64 包含头尾标记和换行符 → 用于 API 请求时必须手动删除；CMD 命令行长度限制导致无法粘贴大段 Base64 → 换用 Python 或 PowerShell 脚本封装。