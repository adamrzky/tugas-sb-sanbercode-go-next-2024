package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Credentials map for Basic Auth
var validCredentials = map[string]string{
	"admin":   "password",
	"editor":  "secret",
	"trainer": "rahasia",
}

// BasicAuth is a middleware for handling Basic Auth
func BasicAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user, pass, ok := r.BasicAuth()
		if !ok || !checkCredentials(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}
		h(w, r, ps)
	}
}

// checkCredentials validates the provided username and password
func checkCredentials(username, password string) bool {
	if pass, ok := validCredentials[username]; ok {
		return pass == password
	}
	return false
}
