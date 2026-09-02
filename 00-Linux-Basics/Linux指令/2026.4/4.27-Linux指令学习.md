今日命令：systemctl、journalctl——systemd 服务管理双核心

```bash
# 1. systemctl — 管理系统服务
systemctl status mini-apiserver          # 查看服务状态
systemctl enable --now mini-apiserver    # 设置开机自启 + 立刻启动
systemctl daemon-reload                  # 修改服务文件后重载配置
systemctl stop/restart mini-apiserver    # 停止/重启服务

# 2. journalctl — 查看 systemd 服务日志
journalctl -u mini-apiserver -f          # 实时跟踪服务日志
journalctl -u mini-apiserver -n 20       # 查看最近20条日志
journalctl -u mini-apiserver --since "10 minutes ago"  # 查看指定时间范围日志
```

底层本质：systemd 替代了传统的 SysV init，通过 cgroup 监控进程状态，Restart=always 本质是利用内核的进程追踪机制。journalctl 读取的是二进制日志 /var/log/journal/，比纯文本日志更高效。

映射之前所学：systemd 和 nohup 是同一需求的两种实现——守护进程管理。nohup 是 Unix 传统方式，systemd 是现代 Linux 标配。
