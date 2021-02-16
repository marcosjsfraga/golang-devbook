package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// Logger writes request info in the terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(" %s %s %s", r.Method, r.RequestURI, r.Host)

		next(w, r)
	}
}

// Authenticate verify if user that made request is authenticated
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("\nAuthenticating...")

		next(w, r)
	}
}
