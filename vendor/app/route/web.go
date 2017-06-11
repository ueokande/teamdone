package route

import (
	"app/controller"
	"net/http"
	"strings"
)

const KeyLen = 32

func Web(c *controller.Context) http.Handler {
	return &webHandler{c: c}
}

type webHandler struct {
	c *controller.Context
}

func (h *webHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	switch {
	case r.URL.Path == "/settings":
		h.c.SettingsGet(w, r)
	case r.URL.Path == "/":
		h.c.HomeGet(w, r)
	case strings.HasPrefix(r.URL.Path, "/") && len(r.URL.Path[1:]) == KeyLen:
		key := r.URL.Path[1:]
		h.c.OrgGet(key, w, r)
	default:
		h.c.NotFound(w, r)
	}
}
