package check

import (
	"audience/conf"
	"testing"
)

var s []conf.HttpServer

func TestHttpCheckSilent(t *testing.T) {

	servers := []conf.HttpServer{
		{Address: "http://localhost:8080", CheckInterval: 10, Tolerance: 5, PushInterval: 500},
		{Address: "http://localhost:8081", CheckInterval: 5, Tolerance: 0, PushInterval: 300},
	}
	for _, v := range servers {
		go HttpCheckSilent(v)
	}
	ch := make(chan int, 0)
	ch <- 1
}
