package route

import (
	"app/controller"
	"net/http"
)

func Api(c *controller.Context) http.Handler {
	return &apiHandler{c: c}
}

type apiHandler struct {
	c *controller.Context
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && r.URL.Path == "/org/create":
		h.c.OrgCreateApi(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/session/get":
		h.c.SessionGetApi(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/session/create":
		h.c.SessionCreateApi(w, r)
	default:
		http.NotFound(w, r)
	}
}
