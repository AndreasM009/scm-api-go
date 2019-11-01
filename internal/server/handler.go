package server

import (
	"encoding/json"
	"net/http"
	"scm-api/internal/domainobjects"
	"scm-api/internal/repository"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func makeCustomerGetHandler(router *mux.Router, repository repository.Repository) *mux.Router {
	handler := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		c := repository.Get(uuid.MustParse(vars["id"]))
		customers := []domainobjects.Customer{c}
		json, _ := json.Marshal(customers)
		w.Write(json)
	}

	router.HandleFunc("/customer/{id}", handler).Methods("GET")

	return router
}

func makeCustomerGetAllHandler(router *mux.Router, repository repository.Repository) *mux.Router {
	handler := func(w http.ResponseWriter, r *http.Request) {
		json, _ := json.Marshal(repository.GetAll())
		w.Write(json)
	}

	router.HandleFunc("/customer", handler).Methods("GET")

	return router
}

func makeCustomerAddHandler(router *mux.Router, repository repository.Repository) *mux.Router {
	handler := func(w http.ResponseWriter, r *http.Request) {
		customer := &domainobjects.Customer{}
		_ = json.NewDecoder(r.Body).Decode(customer)

		customer.Id = uuid.New()
		repository.Add(customer)
	}

	router.HandleFunc("/customer", handler).Methods("POST")
	return router
}

func makeCustomerDeleteAllHandler(router *mux.Router, repository repository.Repository) *mux.Router {
	handler := func(w http.ResponseWriter, r *http.Request) {
		customers := repository.DeleteAll()
		json, _ := json.Marshal(customers)
		w.Write(json)
	}

	router.HandleFunc("/customer", handler).Methods("DELETE")
	return router
}
