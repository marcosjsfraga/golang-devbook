package middlewares

import (
	"api/src/auth"
	"api/src/response"
	"log"
	"net/http"
)

// Logger writes request info in the terminal
func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(" %s %s %s", r.Method, r.RequestURI, r.Host)

		nextFunc(w, r)
	}
}

// Authenticate verify if user that made request is authenticated
func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidarToken(r); err != nil {
			response.Error(w, http.StatusUnauthorized, err)
			return
		}

		nextFunc(w, r)
	}
}
