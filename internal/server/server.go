package server

import (
	"fmt"
	"net/http"
	"scm-api/internal/repository"
	"strconv"

	"github.com/gorilla/mux"
)

// Server ...
type Server struct {
	port   int
	router *mux.Router
}

// CreateServer ...
func CreateServer(port int, r repository.Repository) *Server {
	s := &Server{}
	s.port = port
	s.router = mux.NewRouter()

	makeCustomerGetHandler(s.router, r)
	makeCustomerGetAllHandler(s.router, r)
	makeCustomerAddHandler(s.router, r)
	makeCustomerDeleteAllHandler(s.router, r)

	return s
}

// Start ...
func (s Server) Start() {
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":"+strconv.Itoa(s.port), s.router)
}
