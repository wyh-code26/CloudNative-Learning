复盘 `mini-apiserver` 项目开发过程。

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

为什么真实 K8s 非要依赖 etcd 这种外部存储？
因为它是控制面与业务面的彻底剥离。API Server 本身只做无状态的请求处理和路由分发，所有业务状态下沉到 etcd。这种设计让 API Server 本身可以水平扩展、可以随时重启，而不会丢失集群的任何关键数据。我的版本将“存储”和“处理”耦合在一个进程内，本质上是单体的控制面。

## 与K8s的映射
| 我的实现 | K8s对应组件 | 差异 |
| :--- | :--- | :--- |
| `PodStore` + `map` | etcd | etcd是持久化、分布式的 |
| `handlePods` | API Server路由 | 真实API Server有更复杂的认证/授权/准入 |
| `/healthz` | 健康检查端点 | 功能相同 |

## 面试锚点
    Q: 你的 mini-apiserver 和真实 K8s API Server 的最大区别是什么？
     A: 最大区别是状态的归属。我的版本把状态（Pod数据）放在 API Server 进程的堆内存里，这是一个有状态服务，重启就丢数据，也无法水平扩展。而真正的 K8s API Server 本身是无状态的，所有集群状态都存储在 etcd 里，所以它可以被任意扩展、滚动更新，这是它高可用的基石。
    Q: 你的项目里用了 sync.RWMutex，为什么用读写锁而不用互斥锁？
     A: 因为“读多写少”的业务特点。K8s API Server 的典型场景里，kubectl get 的次数远多于 kubectl create。sync.RWMutex 允许多个读操作同时进入锁保护的临界区，而写操作独占。所以在高并发读的场景下，它的性能远优于 sync.Mutex。我是在理解 K8s 的典型负载模型后，才决定用读写锁来优化性能的。

```