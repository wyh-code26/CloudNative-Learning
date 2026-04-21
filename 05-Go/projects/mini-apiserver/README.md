- [mini-apiserver](./05-Go/projects/mini-apiserver) - 用Go模拟K8s API Server，理解无状态与水平扩展
# Mini API Server (K8s Style)

## 项目简介
这是一个用 Go 语言实现的轻量级 Kubernetes API Server 模拟服务，实现了 Pod 资源的 CRUD 操作和健康检查接口，旨在深入理解 K8s 控制平面的工作原理。

## 访问地址
- HTTPS: `https://api.wuyuhangcn.com`
- HTTP: `http://api.wuyuhangcn.com`（自动重定向到 HTTPS）

## 技术栈
- 后端语言：Go 1.21
- 部署方式：二进制直接运行 / Docker 容器化
- 反向代理：Nginx
- HTTPS 证书：Let's Encrypt (certbot)

## 核心接口
| 方法 | 路径 | 功能 |
| :--- | :--- | :--- |
| GET | `/healthz` | 服务健康检查 |
| GET | `/api/v1/pods` | 获取所有 Pod 列表 |
| POST | `/api/v1/pods` | 创建新的 Pod |

## 本地运行
```bash
go run main.go
curl http://localhost:8081/healthz