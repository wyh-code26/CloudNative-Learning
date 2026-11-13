
#### 2. 踩坑笔记整理

**今日主题**：回顾你昨天**VS Code输入法修复**的全过程。

```markdown
# VS Code snap版输入法冲突修复

## 现象
snap版VS Code在Ubuntu 24.04下使用fcitx5输入中文时，输入过快会漏字母。

## 根本原因
snap应用运行在沙盒中，默认无法完整访问系统输入法框架（fcitx5/ibus）的接口。

## 解决方案
1. 备份配置： 
   `cp -r ~/.config/Code ~/Code_config_backup`
   `cp -r ~/.vscode ~/vscode_extension_backup`
2. 卸载snap版：`sudo snap remove code`
3. 安装官方deb版
4. 恢复配置: 
`cp -r ~/Code_config_backup/* ~/.config/Code/`
`cp -r ~/vscode_extension_backup/* ~/.vscode/`

## 踩坑记录
网上搜到的 `snap connect code:ibus` 命令已过时，新版snap接口名称已变，直接报错`no plug named "ibus"`。

## 底层映射
这和KVM冲突是同一类问题——**软件隔离与权限边界**。snap的沙盒限制了应用能力，K8s的Pod安全策略也是在定义容器的能力边界。

## 对应面试题
**Q: 容器的隔离和虚拟机的隔离有什么区别？**
A: 容器共享宿主机内核，通过namespace隔离进程/网络/文件系统，隔离性弱于虚拟机但开销更小。snap的沙盒本质也是利用AppArmor/seccomp限制应用能力，和容器的安全限制机制相通。