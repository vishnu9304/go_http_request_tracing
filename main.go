package main

import (
	"time"

	"github.com/vishnu9304/go_http_request_tracing/server"
)

func main() {
	go func() {
		s := server.NewServer("1010")
		s.StartServer()
	}()
	time.Sleep(10 * time.Second)
	for i := 0; i < 3; i++ {
		go server.StartLoadTest()
	}
	time.Sleep(2000 * time.Second)
}
