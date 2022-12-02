package main

import (
	"github.com/jfirme-sys/go-auth-docker.git/database"
	"github.com/jfirme-sys/go-auth-docker.git/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
