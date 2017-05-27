package controller

import (
	"encoding/json"
	"net/http"
)

type MessageDto struct {
	Message string
}

func WriteJson(w http.ResponseWriter, data interface{}) {
	content, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(content)
}

func SessionCreate(w http.ResponseWriter, r *http.Request) {
	WriteJson(w, MessageDto{"OK"})
}

func SessionDelete(w http.ResponseWriter, r *http.Request) {
	WriteJson(w, MessageDto{"OK"})
}
