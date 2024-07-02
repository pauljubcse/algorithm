package github.com/pauljubcse/algorithm

import (
	"fmt"
	"net/http"
	"sync"
)

type WeightedRoundRobin struct {
	// backends       []string
	// normalizedWeights        []int
	// currentIndex   uint32
	currentWeight int
	// maxWeight      int
	currentServer backendServer
	backends      Queue[backendServer]
	mu            sync.Mutex
}

type backendServer struct {
	address string
	weight  int
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)

}
func gcdList(weights []int) int {
	gcdValue := weights[0]
	for _, weight := range weights {
		gcdValue = gcd(gcdValue, weight)
	}
	return gcdValue
}

//	func max(weights []int) int {
//		maxValue := weights[0]
//		for _, weight := range weights {
//			if weight > maxValue {
//				maxValue = weight
//			}
//		}
//		return maxValue
//	}
func normalizeWeights(weights []int, gcdWeight int) []int {
	normalizedWeights := make([]int, len(weights))
	for i, weight := range weights {
		normalizedWeights[i] = weight / gcdWeight
	}
	return normalizedWeights
}
func NewWeightedRoundRobin(backends []string, weights []int) *WeightedRoundRobin {
	gcdWeight := gcdList(weights)
	// fmt.Println(gcdWeight)
	normalizedWeights := normalizeWeights(weights, gcdWeight)
	// fmt.Println(normalizedWeights)
	q := NewQueue[backendServer]()
	n := len(backends)
	// fmt.Print(n)
	for i := 0; i < n; i++ {
		q.Enqueue(backendServer{
			address: backends[i],
			weight:  normalizedWeights[i],
		})
	}
	// fmt.Print(q.size)
	var currentServer backendServer
	if !q.IsEmpty() {
		currentServer = q.head.value
	}
	return &WeightedRoundRobin{
		backends: *q,
		// normalizedWeights: normalizedWeights,
		// currentIndex: 0,
		currentWeight: 0,
		currentServer: currentServer,
		// maxWeight: max(normalizedWeights),
	}
}
func (w *WeightedRoundRobin) NextBackend(req *http.Request) string {
	if w.backends.size == 0 {
		return ""
	}
	fmt.Println(req.Header)
	fmt.Println(w.currentWeight)
	w.mu.Lock()
	// fmt.Println(w.currentServer)
	w.currentWeight = w.currentWeight + 1
	if w.currentWeight >= w.currentServer.weight {
		w.currentWeight = 0
		temp, _ := w.backends.Dequeue()
		w.backends.Enqueue(temp)
		w.currentServer, _ = w.backends.Peek()
	}
	w.mu.Unlock()
	return w.currentServer.address
}
