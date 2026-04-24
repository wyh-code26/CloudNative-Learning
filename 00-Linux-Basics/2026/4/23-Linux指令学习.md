```bash
# sort — 排序文本行

sort file.txt                          # 默认字典序升序（a→z，0→9 按字符）
sort -r file.txt                       # 降序排列
sort -n file.txt                       # 按数值排序（9 在 10 前，而非字符序）
sort -t',' -k2 data.csv                # 以逗号分隔，按第2列排序
sort -u file.txt                       # 排序并去重（等效 sort | uniq）

# uniq — 处理相邻重复行

sort data.txt | uniq                   # 去重（必须先行排序）
sort data.txt | uniq -c                # 去重并统计出现次数
sort data.txt | uniq -d                # 仅显示有重复的行（出现≥2次）
sort data.txt | uniq -u                # 仅显示无重复的行（唯一出现）

# 实战组合：分析访问日志 Top 5 IP
sort access.log | cut -d' ' -f1 | uniq -c | sort -rn | head -5
```

底层本质

命令 核心机制 对应 OS 原理
sort 外部归并排序：内存能装下时用堆排序，装不下时自动拆分为临时文件分别排序再合并 虚拟内存管理——超出物理内存时用磁盘做 swap，与 sort 的临时文件策略同源
uniq 逐行状态机：维护“当前行”和“计数”，行变化时输出上一组结果 类似 TCP 滑动窗口——只保留最小必要上下文，不需要回看整个流
` `（管道） 内核匿名管道：前命令 stdout → 内核缓冲区 → 后命令 stdin，文件描述符流转

对应之前所学

· grep 是行过滤，uniq 是相邻行压缩——二者组合可实现复杂文本清洗
· ps aux | grep nginx 中管道的作用，与今天 sort | uniq -c | sort -rn 完全一致

排错预判

症状 原因 解法
10 排在 2 前面 未加 -n，按字符序 加 -n 启用数值比较
uniq 没去重干净 uniq 只看相邻行 先 sort 再 uniq
CSV 排序列不对 分隔符没指定或列号算错 确认 -t 和 -k，列从 1 开始

---