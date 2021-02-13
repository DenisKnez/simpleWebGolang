package middleware

import (
	"fmt"
	"net/http"
)


//NotifyMiddleware is a middleware used to log every incoming request in the console
func NotifyMiddleware() Adapter {
	return func(handler func(w http.ResponseWriter, r *http.Request)) (func(w http.ResponseWriter, r *http.Request)) {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("before")
			handler(w, r)
			fmt.Println("after")
		})
	}
}