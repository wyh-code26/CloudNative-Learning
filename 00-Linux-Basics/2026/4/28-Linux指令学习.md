今日命令：docker update、grep -rn、docker inspect

```bash
# 1. docker update — 修改已有容器的配置（不停机）
docker update --restart=unless-stopped midun-zkp
# --restart=unless-stopped: 容器退出时自动重启，除非手动 docker stop
# 等价选项: no(默认), always, on-failure, unless-stopped

# 2. grep -rn — 递归搜索文件内容，显示行号
grep -rn "midun-dev-key-2026\|BEGIN.*PRIVATE KEY" ~/midun/ 2>/dev/null
# -r: 递归子目录
# -n: 显示行号
# \|: 正则"或"，同时搜索两个模式
# 2>/dev/null: 丢弃权限错误等无关输出

# 3. docker inspect — 查看容器详细配置（验证配置是否生效）
docker inspect midun-zkp | grep -A 3 "RestartPolicy"
```

底层本质：docker update 修改的是容器的 HostConfig，这个配置存储在 /var/lib/docker/containers/<id>/hostconfig.json 里。Docker 引擎重启时会读取这个文件，根据 RestartPolicy 决定是否自动启动容器。--restart=unless-stopped 的逻辑是：只要不是被人手动 docker stop 的，任何情况下退出都自动重启。

对应之前所学：和 systemctl enable --now 一样，docker update --restart=unless-stopped 本质是在容器管理器（Docker Engine / systemd）中注册“这个进程挂了要自动拉起来”的策略。systemd 用 cgroup 监控进程状态，Docker 用 containerd 监控容器状态，底层机制不同但策略等价。

排错预判：docker update 执行成功但重启后发现容器没自动启动 → 用 docker inspect midun-zkp | grep RestartPolicy 确认配置是否已持久化。如果显示 "Name": ""，说明配置没保存成功，重新执行一次即可。
