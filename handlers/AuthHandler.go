package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/DenisKnez/simpleWebGolang/domains"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

//Claims auth claims
type Claims struct {
	jwt.StandardClaims
}

//AuthHandler authentication handler
type AuthHandler struct {
	authUseCase domains.AuthUseCase
	config      viper.Viper
}

//NewAuthHandler creates a new auth handler
func NewAuthHandler(authUseCase domains.AuthUseCase, config viper.Viper) *AuthHandler {
	return &AuthHandler{authUseCase, config}
}

//Refresh refresh the token
func (handler *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {

	bearerToken := r.Header.Get("Authorization")

	if bearerToken == "" {
		http.Error(w, "Required authentication token not provided", http.StatusUnauthorized)
		return
	}

	tokenString := strings.TrimPrefix(bearerToken, "Bearer ")

	if tokenString == "" {
		http.Error(w, "Invalid authentication token", http.StatusUnauthorized)
		return
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(handler.config.GetString("Auth.JwtRefreshTokenSecretKey")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			http.Error(w, "Invalid signature", http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString([]byte(handler.config.GetString("Auth.JwtAccessTokenSecretKey")))

	if err != nil {
		http.Error(w, "Failed to sign the token", http.StatusInternalServerError)
		return
	}

	tokenResponse := &struct {
		AccessToken  string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}

	err = json.NewEncoder(w).Encode(tokenResponse)

	if err != nil {
		http.Error(w, "Failed to encode the resonse", http.StatusInternalServerError)
	}

}

//TestToken used to get a token to test a endpoint
func (handler *AuthHandler) TestToken(w http.ResponseWriter, r *http.Request) {

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "simpleWebGolang",
			ExpiresAt: expirationTime.Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken := jwt.New(jwt.SigningMethodHS384)

	accessTokenString, err := accessToken.SignedString([]byte(handler.config.GetString("Auth.JwtAccessTokenSecretKey")))
	refreshTokenString, err := refreshToken.SignedString([]byte(handler.config.GetString("Auth.JwtRefreshTokenSecretKey")))

	if err != nil {
		http.Error(w, "Failed to sign token into a token string", http.StatusInternalServerError)
		return
	}

	tokenResponse := &struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	err = json.NewEncoder(w).Encode(tokenResponse)

	if err != nil {
		http.Error(w, "Failed to encode the resonse", http.StatusInternalServerError)
	}

	return
}
