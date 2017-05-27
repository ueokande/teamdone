package route

import (
	"fmt"
	"net/http"
)

type ApiHandler struct{}

func (h ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && r.URL.Path == "/session/create":
		fmt.Fprintf(w, "login")
	case r.Method == http.MethodDelete && r.URL.Path == "/session/delete":
		fmt.Fprintf(w, "logout")
	default:
		http.NotFound(w, r)
	}
}
