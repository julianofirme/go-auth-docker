package main

import (
	"github.com/jfirme-sys/go-auth-docker.git/server"
)

func main() {
	server := server.NewServer()

	server.Run()
}
