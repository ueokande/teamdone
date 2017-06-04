package controller

import (
	"net/http"
)

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	jsonError(w, "internal server error", http.StatusInternalServerError)
}
