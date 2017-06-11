package controller

import (
	"encoding/json"
	"net/http"
)

type JsonMessage struct {
	Message string
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	jsonError(w, "internal server error", http.StatusInternalServerError)
}

func jsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
}

func jsonOk(w http.ResponseWriter, data interface{}) error {
	content, err := json.Marshal(data)
	if err != nil {
		return err
	}
	jsonHeader(w)
	w.WriteHeader(http.StatusOK)
	w.Write(content)
	return nil
}

func jsonError(w http.ResponseWriter, message string, code int) {
	content, _ := json.Marshal(JsonMessage{
		Message: message,
	})

	jsonHeader(w)
	w.WriteHeader(code)
	w.Write(content)
}

func (c *Context) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("page not found"))
}

func (c *Context) NotFoundApi(w http.ResponseWriter, r *http.Request) {
	jsonError(w, "page not found", http.StatusNotFound)
}
