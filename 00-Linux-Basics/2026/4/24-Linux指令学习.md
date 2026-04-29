主题：进程管理 ps aux、top、kill

```bash
# 查看所有进程
ps aux

# 实时监控进程
top

# 强制结束进程
kill -9 <PID>
```

笔记要点：

· ps aux 显示所有进程，STAT 列 S 表示休眠，R 表示运行，Z 表示僵尸
· top 按 q 退出，k 键可杀进程
· kill -9 发送 SIGKILL 信号，进程无法捕获，必然终止
· 底层映射：ps 遍历 /proc 文件系统读取内核进程表；kill 通过系统调用向目标进程发送信号

---