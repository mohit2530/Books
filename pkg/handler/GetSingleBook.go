package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mohit2530/Books/pkg/mock"
)

func GetSingleBook(rw http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)                 // reading the dynamic request
	id, err := strconv.Atoi(params["id"]) // converting map to string -> mux returns map

	if err != nil {
		log.Fatal("Internal Server Erorr")
	}

	for _, book := range mock.Books {
		if book.Id == id { // if id are equal
			rw.WriteHeader(http.StatusOK)
			rw.Header().Add("Content-Type", "application-json")
			json.NewEncoder(rw).Encode(book)
			break
		}
	}

}
