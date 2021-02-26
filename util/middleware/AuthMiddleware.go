package middleware

import (
	"net/http"

	"strings"

	"github.com/DenisKnez/simpleWebGolang/domains"
	"github.com/dgrijalva/jwt-go"
	viper "github.com/spf13/viper"
)

//Claims auth claims
type Claims struct {
	jwt.StandardClaims
}

//AuthMiddleware is a middleware authentication
func AuthMiddleware(authUseCase domains.AuthUseCase, config viper.Viper) Adapter {
	return func(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
				return []byte(config.GetString("Auth.JwtAccessTokenSecretKey")), nil
			})

			if err != nil {

				if err == jwt.ErrSignatureInvalid {
					http.Error(w, "Invalid token signature", http.StatusUnauthorized)
					return
				}

				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			handler(w, r)
		})
	}
}
