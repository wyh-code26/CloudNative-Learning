// 声明当前文件属于main包，Go语言中只有main包的程序才能编译为可执行文件，是程序的入口包
package main

// 导入程序所需的标准库依赖包
import (
	"encoding/json" // 处理JSON数据的序列化与反序列化，用于接口请求和响应的JSON格式处理
	"fmt"           // 格式化输入输出，用于控制台打印启动信息、日志等
	"net/http"      // Go原生的HTTP服务库，用于搭建Web接口服务，处理HTTP请求与响应
	"sync"          // 同步原语库，提供读写锁RWMutex，保证并发场景下数据的线程安全
)

// Pod 结构体，模拟Kubernetes中的Pod资源模型，定义了Pod的核心属性
type Pod struct {
	Name      string `json:"name"`      // Pod的名称，json标签指定序列化/反序列化时对应的JSON字段名为name
	Namespace string `json:"namespace"` // Pod所属的命名空间，用于隔离资源，json标签对应字段名为namespace
	Status    string `json:"status"`    // Pod的运行状态，如Running/Pending等，json标签对应字段名为status
}

// PodStore 结构体，模拟K8s的etcd存储组件，实现Pod资源的内存存储与并发安全控制
type PodStore struct {
	mu   sync.RWMutex   // 读写互斥锁，读多写少场景下，允许多个读操作并发，仅写操作互斥，提升并发性能
	pods map[string]Pod // 存储Pod的核心map，key格式为namespace/name，保证全局唯一，value为Pod结构体本身
}

// NewPodStore 构造函数，初始化并返回一个PodStore的指针实例，是PodStore的初始化入口
func NewPodStore() *PodStore {
	// 返回初始化完成的PodStore结构体指针
	return &PodStore{
		pods: make(map[string]Pod), // 初始化存储Pod的map，避免空map赋值引发panic
	}
}

// handlePods PodStore的方法，是/api/v1/pods接口的总入口处理函数，统一分发不同HTTP方法的请求
// 参数w：用于向客户端返回HTTP响应
// 参数r：客户端发来的HTTP请求的所有信息，包括请求方法、请求体、请求头等
func (s *PodStore) handlePods(w http.ResponseWriter, r *http.Request) {
	// 通过switch判断HTTP请求方法，分发到对应的处理函数
	switch r.Method {
	case "GET": // 如果是GET请求，执行查询Pod列表的逻辑
		s.listPods(w, r)
	case "POST": // 如果是POST请求，执行创建Pod的逻辑
		s.createPod(w, r)
	default: // 除了GET、POST之外的其他请求方法，均返回方法不允许的错误
		// 向客户端返回405状态码和错误提示
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// listPods PodStore的方法，处理GET请求，实现查询所有Pod列表的功能
func (s *PodStore) listPods(w http.ResponseWriter, _ *http.Request) {
	s.mu.RLock()         // 加读锁，多个goroutine可以同时加读锁，保证读操作并发安全，阻止写操作，避免读脏数据
	defer s.mu.RUnlock() // defer关键字，确保当前函数退出时，无论是否发生异常，都会释放读锁，避免锁泄漏导致死锁

	// 初始化一个Pod类型的切片，预分配容量为map的长度，减少切片扩容的性能开销
	pods := make([]Pod, 0, len(s.pods))
	// 遍历存储Pod的map，把所有Pod对象逐个追加到切片中
	for _, pod := range s.pods {
		pods = append(pods, pod)
	}

	// 设置HTTP响应头，指定响应体的内容格式为JSON，告知客户端如何解析返回数据
	w.Header().Set("Content-Type", "application/json")
	// 创建JSON编码器，直接把Pod切片序列化为JSON格式，写入HTTP响应体中返回给客户端
	json.NewEncoder(w).Encode(pods)
}

// createPod PodStore的方法，处理POST请求，实现创建Pod的功能
func (s *PodStore) createPod(w http.ResponseWriter, r *http.Request) {
	// 声明一个Pod类型的变量，用于接收并存储客户端请求体中解析出来的Pod数据
	var pod Pod
	// 从HTTP请求体中读取JSON数据，反序列化到上面声明的pod变量中
	// 如果解析失败（如JSON格式错误、字段类型不匹配），则进入错误处理逻辑
	if err := json.NewDecoder(r.Body).Decode(&pod); err != nil {
		// 向客户端返回400状态码（请求参数错误）和具体的错误信息
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // 终止函数执行，不再往下走创建逻辑
	}

	s.mu.Lock()         // 加写锁，同一时间只能有一个goroutine加写锁，保证写操作的原子性，同时阻塞所有读和其他写操作
	defer s.mu.Unlock() // defer关键字，确保函数退出时一定会释放写锁，避免死锁

	// 按照namespace/name的格式拼接Pod的唯一key，和map的key规则保持一致
	key := pod.Namespace + "/" + pod.Name
	// 判断map中是否已经存在该key的Pod，避免重复创建
	if _, exists := s.pods[key]; exists {
		// 如果已存在，向客户端返回409状态码（资源冲突）和错误提示
		http.Error(w, "Pod already exists", http.StatusConflict)
		return // 终止函数执行，不执行后续的写入操作
	}

	// 给新创建的Pod设置默认状态为Running，模拟K8s中Pod创建成功后的运行状态
	pod.Status = "Running"
	// 把处理好的Pod对象存入map中，完成Pod的创建存储
	s.pods[key] = pod

	// 设置HTTP响应头，指定响应体格式为JSON
	w.Header().Set("Content-Type", "application/json")
	// 向客户端返回201状态码（资源创建成功），是RESTful API的规范用法
	w.WriteHeader(http.StatusCreated)
	// 把创建成功的Pod对象序列化为JSON，写入响应体返回给客户端
	json.NewEncoder(w).Encode(pod)
}

// main函数，程序的唯一入口，程序启动时会优先执行main函数
func main() {
	// 调用构造函数，初始化Pod存储实例，后续所有接口操作都基于这个实例
	store := NewPodStore()

	// 注册HTTP路由，把/api/v1/pods路径的请求，交给store的handlePods方法处理
	http.HandleFunc("/api/v1/pods", store.handlePods)

	// 注册健康检查接口/healthz，是K8s中服务健康探测的标准接口，用于判断服务是否正常运行
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		// 向客户端返回200状态码，表示服务健康
		w.WriteHeader(http.StatusOK)
		// 向响应体写入ok字符串，作为健康检查的响应内容
		w.Write([]byte("ok"))
	})

	// 控制台打印服务启动的提示信息，告知用户服务监听的端口
	fmt.Println("Mini API Server (K8s style) listening on :8081")
	// 控制台打印接口说明，列出所有可用的接口
	fmt.Println("Endpoints:")
	fmt.Println("  GET  /api/v1/pods  - list all pods")
	fmt.Println("  POST /api/v1/pods  - create a pod")
	fmt.Println("  GET  /healthz      - health check")

	// 启动HTTP服务，监听0.0.0.0:8081地址，0.0.0.0表示允许所有IP访问，端口为8081
	// 第二个参数nil表示使用Go默认的ServeMux路由管理器
	// 该方法会阻塞运行，直到服务退出
	http.ListenAndServe("0.0.0.0:8081", nil)
}
