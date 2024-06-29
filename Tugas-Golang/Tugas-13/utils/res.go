package utils

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func BasicAuth(handler http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Check basic auth here, and if failed, return error
		user, pass, ok := r.BasicAuth()
		if !ok || !checkCredentials(user, pass) {
			http.Error(w, "Unauthorized.", 401)
			return
		}
		// If authentication is successful, call the handler
		handler(w, r)
	}
}

// Dummy credential check (implement your check here)
func checkCredentials(username, password string) bool {
	// Example: return true if username and password are correct
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
		h(w, r, ps)
	}
}
