package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Pod 模拟K8s的Pod资源
type Pod struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Status    string `json:"status"`
}

// PodStore 模拟etcd的内存存储
type PodStore struct {
	mu   sync.RWMutex
	pods map[string]Pod // key: namespace/name
}

func NewPodStore() *PodStore {
	return &PodStore{
		pods: make(map[string]Pod),
	}
}

// 处理 /api/v1/pods 请求
func (s *PodStore) handlePods(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.listPods(w, r)
	case "POST":
		s.createPod(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *PodStore) listPods(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	pods := make([]Pod, 0, len(s.pods))
	for _, pod := range s.pods {
		pods = append(pods, pod)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pods)
}

func (s *PodStore) createPod(w http.ResponseWriter, r *http.Request) {
	var pod Pod
	if err := json.NewDecoder(r.Body).Decode(&pod); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	key := pod.Namespace + "/" + pod.Name
	if _, exists := s.pods[key]; exists {
		http.Error(w, "Pod already exists", http.StatusConflict)
		return
	}

	pod.Status = "Running"
	s.pods[key] = pod

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pod)
}

func main() {
	store := NewPodStore()

	// 模拟K8s API路由
	http.HandleFunc("/api/v1/pods", store.handlePods)

	// 健康检查端点（水平扩展的关键）
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	fmt.Println("Mini API Server (K8s style) listening on :8081")
	fmt.Println("Endpoints:")
	fmt.Println("  GET  /api/v1/pods  - list all pods")
	fmt.Println("  POST /api/v1/pods  - create a pod")
	fmt.Println("  GET  /healthz      - health check")
	http.ListenAndServe("0.0.0.0:8081", nil)
}
