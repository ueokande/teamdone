package route

import (
	"app/controller"
	"net/http"
	"strings"
)

const KeyLen = 32

type WebHandler struct {
	C *controller.Context
}

func (h WebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	switch {
	case r.URL.Path == "/settings":
		h.C.SettingsGet(w, r)
	case r.URL.Path == "/":
		h.C.HomeGet(w, r)
	case strings.HasPrefix(r.URL.Path, "/") && len(r.URL.Path[1:]) == KeyLen:
		key := r.URL.Path[1:]
		h.C.OrgGet(key, w, r)
	default:
		http.NotFound(w, r)
	}
}
