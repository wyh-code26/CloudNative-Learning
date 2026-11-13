```makedown
命令：ls -lha
 
- 高频场景1：查看当前目录下所有文件（含隐藏）的详细信息  ls -lha 
- 高频场景2：查看指定目录  ls -lha /var/log 
- 踩坑点： -h  必须和  -l  一起用才生效，单独  ls -h  无效
- 底层发生了什么：调用  stat()  系统调用读取每个文件的inode信息（权限、大小、时间戳），按指定格式输出
- C语言联动（10分钟实操）：
写一段C代码，调用  stat()  获取文件大小并打印。c
  
#include <stdio.h>
#include <sys/stat.h>
int main(int argc, char *argv[]) {
    struct stat st;
    if (stat(argv[1], &st) == 0) {
        printf("Size: %ld bytes\n", st.st_size);
    }
    return 0;
}
 
 
编译： gcc -o mystat mystat.c ，运行： ./mystat mystat.c ，对比  ls -l  输出。
 
 
 
命令：mkdir -p
 
- 高频场景：创建多级目录  mkdir -p CloudNative-Learning/02-Kubernetes/YAML-Templates 
- 踩坑点：不加  -p  且父目录不存在时会报错
- 底层发生了什么：调用  mkdir()  系统调用， -p  选项在用户态递归创建每一级目录
 
 
 
命令：cp -r
 
- 高频场景：递归复制整个目录  cp -r /source /dest 
- 踩坑点：不加  -r  无法复制目录；目标目录已存在时，会把源目录复制到目标目录内部，而不是覆盖
- 底层发生了什么：遍历源目录，对每个文件依次调用  open()  +  read()  +  write()  +  close()  完成复制
```