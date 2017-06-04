package route

import (
	"app/controller"
	"net/http"
)

type ApiHandler struct{}

func (h ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && r.URL.Path == "/org/create":
		controller.OrgCreateApi(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/session/get":
		controller.SessionGetApi(w, r)
	default:
		http.NotFound(w, r)
	}
}
