今日指令：icacls — Windows 下的文件权限管理

```bash
# 在 Windows 终端中修复 .pem 文件权限
icacls "C:\Users\17808\.ssh\wuyuhang的云服务器.pem" /inheritance:r /grant:r "%USERNAME%:R"
```

底层本质：Windows 和 Linux 的权限模型完全不同。Linux 用 chmod 400 把文件设为“仅所有者可读”，Windows 则需要用 icacls 来移除继承权限、只赋予当前用户读取权限。SSH 客户端要求密钥文件只能被所有者访问，否则会拒绝连接。

⚠️ 排错预判：Windows 终端复制粘贴命令时，引号可能自动变成中文引号，导致命令执行失败。如果报错，手动重打引号即可。