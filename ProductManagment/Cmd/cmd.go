package cmd

import (
	handlers "apiModules/ProductManagment/Handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Cmd() {

	log.Println("Server Starting...")

	// router oluşturdun muxtan
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/products", handlers.GetProductsHandler).Methods("GET")
	router.HandleFunc("/api/products/{id}", handlers.GetProductHandler).Methods("GET")
	router.HandleFunc("/api/products", handlers.PostProductHandler).Methods("POST")
	router.HandleFunc("/api/products/{id}", handlers.PutProductHandler).Methods("Put")
	router.HandleFunc("/api/products/{id}", handlers.DeleteProductHandler).Methods("Delete")

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	httpServer.ListenAndServe()
	log.Println("Server Ending...")

}
