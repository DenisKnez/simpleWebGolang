package handlers

import (
	"net/http"
)


type AuthHandler struct {
	authUseCase 
}


func(handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	
}