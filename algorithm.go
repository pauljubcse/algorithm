package algorithm

import (
	"net/http"
)

type LoadBalancer interface {
	NextBackend(req *http.Request) string
}
