package github.com/pauljubcse/algorithm

import (
	"net/http"
)

type LoadBalancer interface {
	NextBackend(req *http.Request) string
}
