```markdown
### 命令：cat
- **高频场景**：查看小文件内容 `cat /etc/hosts`；合并文件 `cat a.txt b.txt > c.txt`
- **踩坑点**：对大文件使用`cat`会刷屏，应用`less`替代
- **底层发生了什么**：调用`open()`打开文件，`read()`读取内容，`write()`写入标准输出
- **C语言联动**：
  ```c
  #include <stdio.h>
  int main(int argc, char *argv[]) {
      FILE *fp = fopen(argv[1], "r");
      char buf[1024];
      while (fgets(buf, sizeof(buf), fp)) {
          printf("%s", buf);
      }
      fclose(fp);
      return 0;
  }
```

编译运行：gcc -o mycat mycat.c && ./mycat /etc/hosts

命令：less

· 高频场景：分页查看大文件 less /var/log/syslog
· 常用快捷键：空格翻页，/搜索，q退出
· 底层本质：只加载当前显示的部分到内存，不像cat一次全加载

命令：tail -f

· 高频场景：实时追踪日志 tail -f /var/log/syslog
· 踩坑点：日志滚动后tail -f会失效，需改用tail -F
· 底层发生了什么：使用inotify机制监控文件变化，有新内容时读取并输出
· 面试锚点：Q: tail -f和tail -F的区别？ A: -f监控文件描述符，日志轮转后失效；-F监控文件名，文件重建后自动重新打开。

```