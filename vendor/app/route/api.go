package route

import (
	"app/controller"
	"net/http"
)

type ApiHandler struct {
	C *controller.Context
}

func (h ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && r.URL.Path == "/org/create":
		h.C.OrgCreateApi(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/session/get":
		h.C.SessionGetApi(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/session/create":
		h.C.SessionCreateApi(w, r)
	default:
		http.NotFound(w, r)
	}
}
