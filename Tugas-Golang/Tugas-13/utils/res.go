package utils

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

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

func ResponseJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func HttpRouterToHandlerFunc(h httprouter.Handle) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ps := httprouter.ParamsFromContext(r.Context())
		if ps == nil {
			log.Println("Params are nil")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		h(w, r, ps)
	}
}
