package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mohit2530/Books/pkg/mock"
)

func GetAllBooks(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(mock.Books)
}
