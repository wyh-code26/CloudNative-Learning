今日命令：sed、nginx、dig、scp（今晚排错实战四件套）

```bash
# 1. sed - 流编辑器，不打开文件直接修改文本
sed -i 's/旧内容/新内容/' file.conf         # 替换文件中第一个匹配
sed -i 's/旧内容/新内容/g' file.conf        # 替换文件中所有匹配
sed -n '5,10p' file.conf                    # 查看第 5 到 10 行

# 今晚实战：修复 Nginx 跳转路径
sed -i 's|https://api.wuyuhangcn.com;|https://api.wuyuhangcn.com/healthz;|' /etc/nginx/sites-available/root-redirect

# 2. nginx -t && systemctl reload nginx — 测试配置语法并热重载
nginx -t                                    # 只检查语法，不实际加载
nginx -t && systemctl reload nginx          # 语法正确则重载（今晚高频操作）

# 3. dig - DNS 查询工具，比 nslookup 更详细
dig +short zkp.wuyuhangcn.com               # 只返回解析结果（IP）
dig +short wuyuhangcn.com @ns1.alidns.com   # 指定阿里云权威 DNS 查询

# 4. scp - 在本地和远程服务器之间传文件
scp 本地文件 root@IP:/远程路径               # 上传到服务器
scp root@IP:/远程路径 本地路径               # 从服务器下载
# 今晚实战：同步增量设计文档到本地
scp -i ~/.ssh/wuyuhang的云服务器.pem root@8.210.229.103:/root/midun/docs/INCREMENT-MODULES.md ~/midun/docs/
```

底层本质

命令 核心机制 对应 OS 原理
sed 逐行读入模式空间，执行编辑命令后输出 流式处理，不加载整个文件到内存，大文件安全
nginx -t Nginx 主进程 fork 子进程测试配置，测试完退出 利用进程隔离验证，不影响运行中的服务
systemctl reload 向 Nginx 主进程发送 SIGHUP 信号，主进程重读配置后优雅重启 worker 热重载的核心：先测试新配置，再逐步替换 worker，连接不中断
dig 直接向 DNS 服务器发送 UDP 查询包 应用层协议，默认走 53 端口，比 /etc/hosts 优先级低
scp 基于 SSH 协议传输文件，全程加密 和 ssh 用同一套密钥认证体系

对应之前所学：sed 和 grep 同属文本处理三剑客家族。grep 负责查找，sed 负责替换。今晚你用 sed 修改 Nginx 跳转路径，就是“查找并替换”的标准场景。

⚠️ 排错预判：nginx -t 通过但浏览器仍报错 → 检查是否忘记 systemctl reload（新配置未生效）。scp 报 Permission denied → 检查密钥路径和文件权限（应为 600）。

---