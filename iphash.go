package github.com/pauljubcse/algorithm

import (
	"hash/fnv"
	"log"
	"net/http"
)

type IPHash struct {
	backends []string
}

func NewIPHash(backends []string) *IPHash {
	return &IPHash{
		backends: backends,
	}
}

func (h *IPHash) NextBackend(req *http.Request) string {
	ip := req.RemoteAddr
	log.Println(req.RemoteAddr)
	hash := fnv.New32a()
	hash.Write([]byte(ip))
	index := hash.Sum32() % uint32(len(h.backends))
	return h.backends[index]
}
