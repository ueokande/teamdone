package route

import (
	"app/controller"
	"net/http"
)

type ApiHandler struct{}

func (h ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && r.URL.Path == "/session/create":
		controller.SessionCreate(w, r)
	case r.Method == http.MethodDelete && r.URL.Path == "/session/delete":
		controller.SessionDelete(w, r)
	default:
		http.NotFound(w, r)
	}
}
