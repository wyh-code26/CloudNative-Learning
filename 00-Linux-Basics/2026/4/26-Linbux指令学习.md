今日命令：nohup、&、pkill、pgrep——后台进程的“生”与“死”

```bash
# 1. nohup + & —— 后台运行，不挂断
nohup /root/mini-apiserver > /var/log/mini-apiserver.log 2>&1 &
# nohup: 免疫 HUP 信号，即使终端关闭，进程继续运行
# &: 后台执行，终端立即返回（可以继续输入其他命令）
# >: 重定向标准输出到日志文件
# 2>&1: 把标准错误也重定向到同一个日志文件

# 2. pgrep —— 按进程名搜索 PID
pgrep -f mini-apiserver          # -f: 匹配完整命令行（包括参数）
# 返回 PID (如: 4257)

# 3. pkill —— 按进程名结束进程
pkill -f mini-apiserver          # 发送 SIGTERM (15)，等待程序自行退出
sleep 1                          # 等待1秒
# 如果进程仍存在，强制结束：
kill -9 $(pgrep -f mini-apiserver)  # 发送 SIGKILL (9)，内核强制终止
```

底层本质

概念 核心机制 对应 OS 原理
nohup 修改进程对 SIGHUP 信号的响应方式为“忽略”，同时自动重定向输出 信号是 Unix/Linux 进程间通信的核心机制，SIGHUP(1) 在终端关闭时自动发送
& Shell 把任务放入后台作业列表，进程与终端分离，但输出默认仍指向终端 Shell 作业控制，jobs 命令可查看所有后台任务
pkill 遍历 /proc 文件系统，找到匹配进程，发送指定信号 信号发送由内核 kill() 系统调用完成
SIGTERM vs SIGKILL SIGTERM 可被程序捕获（优雅退出），SIGKILL 不可捕获（强制终止） 符合 Unix“先礼后兵”的设计哲学

对应之前所学：nohup 守护进程与 systemd 管理服务是同一种需求的两种实现，前者是传统 Unix 方式，后者是现代化守护进程管理器。

排错预判：nohup 启动后进程从未出现 → 检查启动命令是否有语法错误（nohup 会把 stderr 重定向，但不主动显示）。程序在 nohup 下运行但意外退出 → 查看日志文件，通常是程序自己崩溃了，与 nohup 无关。