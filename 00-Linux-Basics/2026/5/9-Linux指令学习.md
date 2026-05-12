今日指令：dd — 磁盘底层操作利器

```bash
# 1. 用 dd 彻底清空 U 盘分区表
sudo dd if=/dev/zero of=/dev/sdb bs=1M count=10

# 2. 用 dd 制作启动盘（虽然对Windows ISO不好用，但对Linux ISO最稳）
sudo dd if=ubuntu-24.04.iso of=/dev/sdb bs=4M status=progress

# 3. 用 dd 备份整个磁盘（你没用上但应该知道）
sudo dd if=/dev/nvme0n1 of=/path/to/backup.img bs=4M status=progress
```

底层本质：dd 是Unix/Linux中最原始的底层数据流处理工具，它不做任何文件系统层面的理解，直接逐字节复制。if 是输入文件（或设备），of 是输出文件（或设备），bs 是每次读写的块大小。它不像cp那样只复制文件，而是直接操作设备节点（如 /dev/sdb），所以能清空分区表、制作启动盘、备份整个硬盘。

对应之前所学：dd 和 gparted/parted 同属磁盘管理工具，但 dd 更底层、更暴力、出错后也更难恢复。执行时务必再三确认 of 后面的设备名，选错盘会导致数据永久丢失。

⚠️ 排错预判：U盘插上后 lsblk 看不到 → 检查USB接口，或者执行 dmesg | tail -20 查看内核日志