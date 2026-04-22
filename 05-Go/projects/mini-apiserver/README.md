markdown
  
# Mini API Server (K8s Style)

## 项目简介
这是一个用 Go 语言实现的轻量级 Kubernetes API Server 模拟服务，实现了 Pod 资源的 CRUD 操作和健康检查接口，旨在深入理解 K8s 控制平面的工作原理。
本服务为纯API后端服务，无前端页面，已完成生产级部署，核心接口可通过公网HTTPS正常访问。

## 访问地址
- HTTPS: `https://api.wuyuhangcn.com`
- HTTP: `http://api.wuyuhangcn.com`（自动重定向到 HTTPS）
> 说明：域名根路径无前端页面，直接访问会返回解析提示，核心业务接口均正常可用

## 技术栈
- 后端语言：Go 1.21
- 部署方式：二进制直接运行 / Docker 容器化
- 反向代理：Nginx
- HTTPS 证书：Let's Encrypt (certbot，已配置自动续期)
- 服务器环境：Ubuntu 24.04 LTS 阿里云ECS

## 核心接口
| 方法 | 路径 | 功能 |
| :--- | :--- | :--- |
| GET | `/healthz` | 服务健康检查 |
| GET | `/api/v1/pods` | 获取所有 Pod 列表 |
| POST | `/api/v1/pods` | 创建新的 Pod |

## 本地运行
```bash
# 启动服务
go run main.go
# 验证健康检查接口
curl http://localhost:8081/healthz
 
 
架构设计
 
本项目模拟了 Kubernetes API Server 的核心设计模式：
 
mermaid
  
graph TD
    subgraph "客户端层"
        A[curl / Web 客户端]
    end
    
    subgraph "接入层"
        B[Nginx 反向代理<br/>HTTPS 终端]
    end
    
    subgraph "应用层"
        C[mini-apiserver<br/>Go HTTP 服务]
        D[PodStore<br/>内存存储 + 读写锁]
    end
    
    A -->|HTTPS 请求| B
    B -->|proxy_pass| C
    C -->|读写操作| D
 
 
设计映射：
 
- PodStore + map → 对应 K8s 的 etcd，负责状态存储
- handlePods 路由分发 → 对应 K8s API Server 路由层
- sync.RWMutex → 实现高并发读写安全，读多写少场景性能更优
-  /healthz  健康检查 → 无状态设计，支持水平扩展
- Nginx 反向代理 → 接入层与业务层解耦，实现负载均衡
 
部署指南
 
Docker 部署
 
bash
  
# 1. 构建镜像
docker build -t mini-apiserver .

# 2. 运行容器
docker run -d --name mini-apiserver -p 8081:8081 mini-apiserver

# 3. 验证服务
curl http://localhost:8081/healthz
 
 
生产环境部署（阿里云 + Nginx + HTTPS）
 
bash
  
# 1. 交叉编译Linux环境二进制文件
GOOS=linux GOARCH=amd64 go build -o mini-apiserver main.go
# 2. 上传二进制到云服务器
scp -i ~/.ssh/...... mini-apiserver root@8.210.229.103:/opt/

# 3. 服务器后台启动服务
ssh -i ~/.ssh/...... root@8.210.229.103
nohup /opt/mini-apiserver > /var/log/mini-apiserver.log 2>&1 &

# 4. Nginx 反向代理配置
# 见项目根目录 nginx.conf 示例
 
 
API 使用示例
 
健康检查
 
bash
  
curl https://api.wuyuhangcn.com/healthz
# 响应: ok
 
 
 创建一个 Pod
 
bash
  
curl -X POST https://api.wuyuhangcn.com/api/v1/pods \
  -H "Content-Type: application/json" \
  -d '{"name":"nginx","namespace":"default"}'
# 响应: {"name":"nginx","namespace":"default","status":"Running"}


列出所有 Pod
 
bash
  
curl https://api.wuyuhangcn.com/api/v1/pods
# 响应: [{"name":"test-pod","namespace":"","status":"Running"}]
 
 

 
 
技术亮点与学习笔记
 
1. 并发安全设计
 
go
  
type PodStore struct {
    mu   sync.RWMutex  // 读写锁
    pods map[string]Pod
}
 
 
- 读操作使用 RLock()，支持高并发读取
- 写操作使用 Lock()，保证数据一致性
- 读多写少场景下性能远优于普通互斥锁
 
2. RESTful API 设计
 
go
  
func (s *PodStore) handlePods(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":  s.listPods(w, r)
    case "POST": s.createPod(w, r)
    }
}
 
 
- 基于 HTTP 方法实现资源操作分发
- 结构清晰，易于扩展更多资源类型（Deployment、Service 等）
 
3. 无状态设计
 
-  /healthz  无状态、无依赖，便于监控与扩容
- 与 K8s 原生 API Server 设计思想一致：业务无状态，状态下沉到存储
 
4. 容器化与多阶段构建
 
dockerfile
  
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o mini-apiserver main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/mini-apiserver .
EXPOSE 8081
CMD ["./mini-apiserver"]
 
 
- 构建与运行分离，大幅减小镜像体积
- 最终镜像仅包含二进制，安全、轻量、启动快
 
5. 生产级部署能力
 
- Nginx 负责 HTTPS 终端、请求转发与日志
- Let's Encrypt 证书自动续期，保证服务长期可用
- 后台守护运行，支持服务持久化稳定运行
- 全链路日志可追溯，便于问题排查
 
待改进项
 
问题 当前实现 改进方向 K8s 对应方案 
状态存储 内存 map，服务重启数据丢失 引入持久化存储（BoltDB/SQLite） etcd 
水平扩展 内存状态不共享，无法多实例扩容 引入外部共享存储 etcd + Watch 机制 
认证授权 无身份认证与权限控制 添加 JWT / RBAC 权限控制 RBAC + Admission Control 
资源版本 无资源版本控制 增加 ResourceVersion 版本机制 etcd ModRevision 
Watch 机制 无事件推送能力 实现长连接 / SSE 事件推送 etcd Watch
  

