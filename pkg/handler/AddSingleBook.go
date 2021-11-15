package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/mohit2530/Books/pkg/mock"
	"github.com/mohit2530/Books/pkg/model"
)

// AddSingleBook method to the database
func AddSingleBook(rw http.ResponseWriter, r *http.Request) {

	defer r.Body.Close() // close the body when request is retrieved
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Internal Server Erorr")
	}
	var book model.Book
	json.Unmarshal(body, &book)
	book.Id = rand.Intn(1000)
	mock.Books = append(mock.Books, book)

	// response
	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("Content-Type", "application-json")
	json.NewEncoder(rw).Encode(book.Id)
}
