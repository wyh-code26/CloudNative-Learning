今日指令：curl 与 Shell 引号规则 — ' 与 " 的本质区别

今天手打多行 curl 命令时，你发现了 -H 和参数值之间没有空格会报错，也顺带理清了 Shell 引号的底层规则。

```bash
# 1. 单引号 '' — 字面量保险箱，内容原样不动
echo '$HOME'                       # 输出 $HOME，不是 /home/wuyuhang
curl -d '{"name":"deom-pod"}'      # JSON 用单引号，内部引号无需转义

# 2. 双引号 "" — 变量和命令先替换，再传给命令
echo "$HOME"                       # 输出 /home/wuyuhang
curl -d "{\"name\":\"deom-pod\"}"  # JSON 用双引号，内部引号必须转义

# 3. 多行续行 — \ 后面必须直接接回车，不能有空格
curl -X POST https://api.wuyuhangcn.com/api/v1/pods \
  -H "Content-Type: application/json" \     # 正确：\ 后直接回车
  -d '{"name":"deom-pod"}'                  # 正确：选项和参数之间有空格
```

底层本质：Shell 的引号规则和续行规则是两套独立机制。引号控制内容如何被解释（原样 vs 替换），续行符 \ 控制物理行如何被拼接成逻辑行。你今天的错误——把 -H 和它的参数值连在一起写——本质上是因为 \ 把换行符吃掉了，导致 -H"..." 被 Shell 当成了一个连续的字符串，而不是一个选项加一个参数。Space 在 Shell 里是参数分隔符，不是可有可无的空格。

映射到《密盾》项目：谢海峰前端 fetch 请求的请求体也是 JSON。JavaScript 里用 JSON.stringify() 自动处理引号转义，而你在终端用 curl 测试时，用单引号包裹 JSON 是最简方式——这就是两种语言对同一个问题的不同解法。

⚠️ 排错预判：curl -d 返回 invalid request body → 90% 是 JSON 引号被 Shell 提前吃掉。解决：JSON 字符串用单引号包裹。