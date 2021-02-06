package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/DenisKnez/simpleWebGolang/data"
	"github.com/DenisKnez/simpleWebGolang/domains"
	"github.com/google/uuid"
)

//BookHandler book handler
type BookHandler struct {
	usecase domains.BookUseCase
}

//NewBookHandler creates a new book handler
func NewBookHandler(usecase domains.BookUseCase) *BookHandler {
	return &BookHandler{usecase}
}

//CreateBook create a new book
func (handler *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		http.Error(w, "Request has no body", http.StatusBadRequest)
		return
	}

	book := data.Book{}
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to decode the body", http.StatusBadRequest)
		return
	}

	err = handler.usecase.CreateBook(&book)

	if err != nil {
		http.Error(w, "Failed to create a new book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return

}

//Books get all books
func (handler *BookHandler) Books(w http.ResponseWriter, r *http.Request) {

	books, err := handler.usecase.Books()

	if err != nil {
		http.Error(w, "Could not return books", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(books)

	if err != nil {
		http.Error(w, "Could not encode books", http.StatusInternalServerError)
	}

	return
}

//PagedBooks get books by the provided parameters
func (handler *BookHandler) PagedBooks(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	pageSizeString := query.Get("pagesize")
	pageNumberString := query.Get("pagenumber")
	if pageSizeString == "" || pageNumberString == "" {
		http.Error(w, "Either pagesize or pagenumber parameter not provided", http.StatusBadRequest)
		return
	}

	pageSize, sizeErr := strconv.Atoi(pageSizeString)
	pageNumber, numErr := strconv.Atoi(pageNumberString)

	if sizeErr != nil || numErr != nil {
		http.Error(w, "Parameters are not numbers", http.StatusBadRequest)
		return
	}

	books, err := handler.usecase.PagedBooks(pageSize, pageNumber)

	if err != nil {
		http.Error(w, "Could not return books", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(books)

	if err != nil {
		http.Error(w, "Could not encode books", http.StatusInternalServerError)
	}

	return
}

//GetBook get book with the provided id
func (handler *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "the id parameter was not provided", http.StatusBadRequest)
		return
	}

	book, err := handler.usecase.GetBookByID(id)

	if err != nil {
		http.Error(w, "Could not find book", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(book)

	if err != nil {
		http.Error(w, "Could not encode book", http.StatusInternalServerError)
	}

	return
}

//DeleteBook delete book with the provided id
func (handler *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Required parameter id was not provided", http.StatusBadRequest)
		return
	}

	err := handler.usecase.DeleteBook(id)

	if err != nil {
		http.Error(w, "Could not delete book", http.StatusInternalServerError)
		return
	}

	return
}

//UpdateBook gets all the Books
func (handler *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Required parameter id was not provided", http.StatusBadRequest)
		return
	}

	Book := data.Book{}
	err := json.NewDecoder(r.Body).Decode(&Book)

	if err != nil {
		http.Error(w, "Failed to decode provided json", http.StatusBadRequest)
		return
	}

	Book.ID, err = uuid.Parse(id)

	if err != nil {
		http.Error(w, "Provided parameter id was not valid", http.StatusBadRequest)
		return
	}

	err = handler.usecase.UpdateBook(&Book)

	if err != nil {
		http.Error(w, "Could not update Book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
