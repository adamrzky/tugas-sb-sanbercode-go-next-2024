package utils

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// BasicAuth is a middleware function that provides basic authentication
func BasicAuth(handler http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user, pass, ok := r.BasicAuth()
		if !ok || !checkCredentials(user, pass) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), httprouter.ParamsKey, ps)
		handler(w, r.WithContext(ctx))
	}
}

// Dummy credential check (implement your check here)
func checkCredentials(username, password string) bool {
	return username == "admin" && password == "password"
}

// ResponseJSON sends a JSON response with a given status code
func ResponseJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// HttpRouterToHandlerFunc converts a httprouter.Handle to http.HandlerFunc
func HttpRouterToHandlerFunc(h httprouter.Handle) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ps := httprouter.ParamsFromContext(r.Context())
		if ps == nil {
			// log.Println("Params are nil. Setting empty Params.")
			ps = httprouter.Params{}
		}
		h(w, r, ps)
	}
}
