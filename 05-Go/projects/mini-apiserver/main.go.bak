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
	pods map[string]Pod
}

func NewPodStore() *PodStore {
	return &PodStore{
		pods: make(map[string]Pod),
	}
}

// handlePods 处理 /api/v1/pods 请求，分发不同方法
func (s *PodStore) handlePods(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.listPods(w, r)
	case "POST":
		s.createPod(w, r)
	case "DELETE":
		s.deletePod(w, r)
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Method not allowed",
		})
	}
}

// listPods 处理 GET 请求，列出所有 Pod
func (s *PodStore) listPods(w http.ResponseWriter, _ *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	pods := make([]Pod, 0, len(s.pods))
	for _, pod := range s.pods {
		pods = append(pods, pod)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pods)
}

// createPod 处理 POST 请求，创建 Pod
func (s *PodStore) createPod(w http.ResponseWriter, r *http.Request) {
	var pod Pod
	if err := json.NewDecoder(r.Body).Decode(&pod); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid JSON format",
		})
		return
	}

	// 参数校验：name 必填
	if pod.Name == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "'name' is required",
		})
		return
	}
	// 参数校验：namespace 必填
	if pod.Namespace == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "'namespace' is required",
		})
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	key := pod.Namespace + "/" + pod.Name
	if _, exists := s.pods[key]; exists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Pod already exists",
		})
		return
	}

	pod.Status = "Running"
	s.pods[key] = pod

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pod)
}

// deletePod 处理 DELETE 请求，删除指定 Pod
func (s *PodStore) deletePod(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	namespace := r.URL.Query().Get("namespace")
	if name == "" || namespace == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "'name' and 'namespace' are required",
		})
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	key := namespace + "/" + name
	if _, exists := s.pods[key]; !exists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Pod not found",
		})
		return
	}

	delete(s.pods, key)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "deleted",
	})
}

// main 程序入口
func main() {
	store := NewPodStore()

	http.HandleFunc("/api/v1/pods", store.handlePods)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	fmt.Println("Mini API Server (K8s style) listening on :8081")
	fmt.Println("Endpoints:")
	fmt.Println("  GET    /api/v1/pods  - list all pods")
	fmt.Println("  POST   /api/v1/pods  - create a pod")
	fmt.Println("  DELETE /api/v1/pods  - delete a pod")
	fmt.Println("  GET    /healthz      - health check")
	http.ListenAndServe("0.0.0.0:8081", nil)
}
