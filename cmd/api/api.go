package main

import (
	"scm-api/internal/inmemrepository"
	"scm-api/internal/server"
)

func main() {
	r := inmemrepository.CreateRepository()
	s := server.CreateServer(8080, r)
	s.Start()
}
