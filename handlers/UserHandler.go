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

//UserHandler user handler
type UserHandler struct {
	userUseCase domains.UserUseCase
}

//NewUserHandler new user handler
func NewUserHandler(userUseCase domains.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase}
}

//GetUser gets the user with the provided id
func (handler *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "user id was not provided", http.StatusBadRequest)
		return
	}

	user, err := handler.userUseCase.GetUserByID(id)

	if err != nil {
		http.Error(w, "The user with the id does not exist", http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		http.Error(w, "Json encoding failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusFound)
	return
}

//GetUsers gets all the users
func (handler *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := handler.userUseCase.Users()

	if err != nil {
		http.Error(w, "Could not return users", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(users)

	if err != nil {
		http.Error(w, "Json encoding failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusFound)
	return
}

//CreateUser create a user
func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Post method was not used", http.StatusBadRequest)
		return
	}

	user := data.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Failed to decode the json provided", http.StatusBadRequest)
		return
	}

	err = handler.userUseCase.CreateUser(&user)

	if err != nil {
		http.Error(w, "Failed to create new user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

//PagedUsers gets users by paged parameters
func (handler *UserHandler) PagedUsers(w http.ResponseWriter, r *http.Request) {

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

	users, err := handler.userUseCase.PagedUsers(pageSize, pageNumber)

	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	fmt.Print(users)

	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusFound)
	return
}

//DeleteUser deletes a user with the provided id
func (handler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Required parameter id was not provided", http.StatusBadRequest)
	}

	err := handler.userUseCase.DeleteUser(id)

	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

//UpdateUser gets all the users
func (handler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Required parameter id was not provided", http.StatusBadRequest)
		return
	}

	user := data.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Failed to decode provided json", http.StatusBadRequest)
		return
	}

	user.ID, err = uuid.Parse(id)

	if err != nil {
		http.Error(w, "Provided parameter id was not valid", http.StatusBadRequest)
		return
	}

	err = handler.userUseCase.UpdateUser(&user)

	if err != nil {
		http.Error(w, "Could not update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
