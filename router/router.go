package router

import (
	"github.com/vatsal278/go-project/controler"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.StartPage).Methods("GET")
	router.HandleFunc("/employee/{id}", controller.DisplayEmployee).Methods("GET")
	router.HandleFunc("/employee", controller.CreateEmployee).Methods("POST")
	router.HandleFunc("/employee/{id}", controller.DeleteEmployee).Methods("DELETE")
	router.HandleFunc("/employee/{id}", controller.UpdateEmployee).Methods("PUT")
	return router
}
