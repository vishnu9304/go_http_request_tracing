package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type InitServer struct {
	port string
}

func NewServer(port string) InitServer {
	server := InitServer{port: port}
	return server
}

func (r InitServer) StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler)
	log.Fatal(http.ListenAndServe(":"+r.port, router))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	time.Sleep(30 * time.Second)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	indexTemplate := `{
	endPoint: "%v",
	method: "%v",
	latency: "%v",
	status: "alive"
}`
	end := time.Since(start)
	res := fmt.Sprintf(indexTemplate, r.URL, r.Method, end)
	io.WriteString(w, res)
}

func StartLoadTest() {
	resp, err := http.Get("http://localhost:1010/")
	if err != nil {
		panic(fmt.Sprintf("Got error: %v", err))
	}
	io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
}
