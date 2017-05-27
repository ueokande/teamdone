package route

import (
	"fmt"
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
		fmt.Fprintf(w, "settings")
	case r.URL.Path == "/":
		fmt.Fprintf(w, "home")
	case strings.HasPrefix(r.URL.Path, "/") && len(r.URL.Path[1:]) == KeyLen:
		fmt.Fprintf(w, "team:"+r.URL.Path[1:])
	default:
		http.NotFound(w, r)
	}
}
