```markdown
### 命令：head
- **高频场景**：查看文件前N行 `head -n 20 /var/log/syslog`
- **踩坑点**：默认是10行，`-n` 参数不可省略数字
- **底层发生了什么**：打开文件，只读取前N行后即关闭，不会加载整个文件
- **C语言联动**：
 
  #include <stdio.h>
  int main(int argc, char *argv[]) {
      FILE *fp = fopen(argv[1], "r");
      char buf[1024];
      int n = atoi(argv[2]);
      for (int i = 0; i < n && fgets(buf, sizeof(buf), fp); i++) {
          printf("%s", buf);
      }
      fclose(fp);
      return 0;
  }


编译：gcc -o myhead myhead.c，运行：./myhead /etc/passwd 5

命令：grep

  高频场景：在文件中搜索关键词 grep -rn "error" /var/log/
  踩坑点：正则表达式中的特殊字符（如 . *）需转义
  底层发生了什么：逐行读取文件，用正则引擎匹配，匹配成功则输出
  面试锚点：Q: grep 和 find 的区别？ A: grep 在文件内容中搜索，find 在文件名和属性中搜索。

命令：wc

  高频场景：统计代码行数 wc -l *.go
  常用组合：ls | wc -l 统计文件个数
  底层本质：读取文件时计数，-l 统计换行符数量

```