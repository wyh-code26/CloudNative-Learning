**今日主题**：复盘 `mini-apiserver` 项目开发过程。

**写作框架**：
```markdown
# mini-apiserver 开发复盘

## 目标
用Go语言模拟K8s API Server的核心行为：处理HTTP请求、管理资源、支持水平扩展。

## 技术点拆解
- HTTP服务：`net/http` 包，路由注册
- 资源管理：用 `map` 模拟 etcd 存储
- 并发安全：`sync.RWMutex` 读写锁
- RESTful API：GET/POST 方法分发
- 健康检查：`/healthz` 端点

## 遇到的坑
- `json.NewDecoder(r.Body).Decode(&pod)` 后需要调用 `r.Body.Close()`
- 多实例共享状态问题：当前内存存储无法真正水平扩展，需引入外部存储

## 根本原因
HTTP服务是无状态的，但状态存储在内存中导致无法跨实例共享。这正是真实K8s API Server依赖etcd的原因。

## 与K8s的映射
| 我的实现 | K8s对应组件 | 差异 |
| :--- | :--- | :--- |
| `PodStore` + `map` | etcd | etcd是持久化、分布式的 |
| `handlePods` | API Server路由 | 真实API Server有更复杂的认证/授权/准入 |
| `/healthz` | 健康检查端点 | 功能相同 |

## 面试锚点
**Q: 你写的mini-apiserver和真实K8s API Server的最大区别是什么？**
A: 状态存储。我的版本用内存map，重启数据丢失，且无法多实例共享。真实API Server用etcd实现持久化和分布式一致性。
```