package middleware

import (
	"net/http"
)

//Adapter type that is used to combine all he middlewares
type Adapter func(func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request)


//Adapt combine all the adapters
func Adapt(handler func(w http.ResponseWriter, r *http.Request), adapters ...Adapter) func(w http.ResponseWriter, r *http.Request) {
	for _, adapter := range adapters {
		handler = adapter(handler)
	}
	return handler
}