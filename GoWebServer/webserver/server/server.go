package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Port int
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TestHandler")
}

func (s *Server) Start() {
	fmt.Println("hello from server", s.Port)

	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/test/", handleTest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
