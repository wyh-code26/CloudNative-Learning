```markdown
# Host-Only网路配置与静态IP地址设置

## 现象
- 初始状态：虚拟机Host-Onl网卡通过DHCP获取IP（192.168.56.101）
- 期望状态：固定IP为192.168.56.10

## 操作步骤
1. 编辑 `/etc/netplan/00-install-config.yaml`
2. 设置 `dhcp4: no` 和 `address: [192.168.56.10/24]`
3. `sudo nerplan apply`

## 遇到的坑
- 应用配置后仍有DHCP残留IP(192.168.56.101)
- Netplan配置文件警告(`Permissions are too open`)

## 解决方案
- 释放DHCP租约：`sudo dhclient -r enp0s8`
- 修复权限:`sudo chomd 600 /etc/netplan/00-installer-config.yaml`

## 根本原因一句话
DHCP租约未释放，导致静态IP与动态IP共存；Netplan要求配置文件仅root可读写以保证安全。

## 底层映射
这本质上是“固定端点 vs 动态发现”的选择问题。k8s Service的ClusterIP就是静态端点，Pod IP是动态分配。

## 对应面试题
**Q: k8s Service的ClusterIP为什么是固定的？Pod重启后IP会变，Service怎么找到新的Pod？**
A：Service通过Label Selector和Endpoints对象维护后端Pod的实时IP列表。ClusterIP固定是为了给集群内提供一个稳定的访问入口。 (这和虚拟机摄静态IP是一个道理)
