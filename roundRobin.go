package github.com/pauljubcse/algorithm

import (
	"net/http"
	"sync/atomic"
)

type RoundRobin struct {
	backends []string
	current  uint32
}

func wrapAround(r *RoundRobin) { //No need for using mutext to implement %, this is just doen to prevent opverflow in long term
	r.current = r.current % uint32(len(r.backends))
}
func NewRoundRobin(backends []string) *RoundRobin {
	return &RoundRobin{
		backends: backends,
	}
}

func (r *RoundRobin) NextBackend(req *http.Request) string { //passinmg request to maintain consistent interface, here request is not required, but will be required in case of ip_hash
	index := atomic.AddUint32(&r.current, 1)
	go wrapAround(r)
	return r.backends[index%uint32(len(r.backends))]
}
