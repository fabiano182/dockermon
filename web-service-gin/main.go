package main

import (
	"github.com/fabiano182/dockermon/web-service-gin/server"
)

func main() {
	s := server.NewServer()

	s.Run()
}
