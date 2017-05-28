package route

import (
	"app/controller"
	"net/http"
	"strings"
)

const KeyLen = 32

type WebHandler struct{}

func (h WebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	switch {
	case r.URL.Path == "/settings":
		controller.SettingsGet(w, r)
	case r.URL.Path == "/":
		controller.HomeGet(w, r)
	case strings.HasPrefix(r.URL.Path, "/") && len(r.URL.Path[1:]) == KeyLen:
		key := r.URL.Path[1:]
		controller.OrgGet(key, w, r)
	default:
		http.NotFound(w, r)
	}
}
