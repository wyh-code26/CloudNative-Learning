今日指令：fsck、grub-install、update-grub — 系统急救三件套

```bash
# 1. fsck — 检查和修复文件系统
sudo fsck -y /dev/nvme0n1p2     # -y 表示对全部错误自动回答 yes

# 2. grub-install — 重新安装 GRUB 引导加载器
sudo grub-install /dev/nvme0n1

# 3. update-grub — 更新 GRUB 配置文件
sudo update-grub
```

底层本质：fsck 检查的是文件系统的元数据结构（超级块、inode表、空闲块位图），如果这些元数据因为非法关机而损坏，fsck 会尝试修复它们。grub-install 把一个小的引导镜像写入磁盘的最开头（MBR或EFI分区），让BIOS/UEFI在开机时能找到系统。update-grub 扫描硬盘上所有已安装的操作系统，自动生成开机菜单配置文件 /boot/grub/grub.cfg。

对应之前所学：你之前在Live USB里多次尝试挂载分区失败，就是因为文件系统元数据已经损坏到连 fsck 都无法修复的程度。这是强制关机对固态硬盘最严重的影响——不是硬件损坏，而是元数据逻辑结构崩溃。

⚠️ 排错预判：grub-install 提示找不到EFI分区 → 确认 /boot/efi 已挂载，执行 lsblk 确认 EFI 分区位置