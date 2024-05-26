package main

import (
	"fmt"

	"example.com/webserver/server"
)

func main() {
	fmt.Println("main")
	s := server.Server{Port: 8080}
	s.Start()
}
