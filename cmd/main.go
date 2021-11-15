package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohit2530/Books/pkg/handler"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", handler.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handler.GetSingleBook).Methods(http.MethodGet)
	router.HandleFunc("/books", handler.AddSingleBook).Methods(http.MethodPost)

	log.Println("Api is up and running ... ")
	http.ListenAndServe(":4000", router)
}
