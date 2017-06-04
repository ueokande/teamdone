package controller

import (
	"app/model"
	"app/session"
	"database/sql"
	"encoding/json"
	"net/http"
)

type SessionGetApiDto struct {
	Id   int64
	Name string
}

type SessionCreateApiForm struct {
	Name string
}

type SessionCreateApiDto struct {
	Id   int64
	Name string
}

func SessionGetApi(w http.ResponseWriter, r *http.Request) {
	s, err := session.DefaultSessionManager().StartSession(w, r)
	if err != nil {
		InternalServerError(w, r)
		return
	}
	uidf, ok := s.Values["user_id"].(float64)
	uid := int64(uidf)
	if !ok {
		jsonError(w, "session not found", http.StatusNotFound)
		return
	}
	u, err := model.UserById(uid)
	if err == sql.ErrNoRows {
		jsonError(w, "session not found", http.StatusNotFound)
		return
	} else if err != nil {
		InternalServerError(w, r)
		return
	}

	jsonOk(w, SessionGetApiDto{
		Id:   u.Id,
		Name: u.Name,
	})
}

func SessionCreateApi(w http.ResponseWriter, r *http.Request) {
	s, err := session.DefaultSessionManager().StartSession(w, r)
	if err != nil {
		InternalServerError(w, r)
		return
	}
	uidf, ok := s.Values["user_id"].(float64)
	uid := int64(uidf)
	if ok {
		jsonError(w, "already created", http.StatusBadRequest)
		return
	}

	var form SessionCreateApiForm
	err = json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		jsonError(w, "invalid json", http.StatusBadRequest)
		return
	}
	if len(form.Name) == 0 {
		jsonError(w, "user name is required", http.StatusBadRequest)
		return
	}

	uid, err = model.UserCreate(form.Name)
	if err != nil {
		InternalServerError(w, r)
		return
	}
	jsonOk(w, SessionGetApiDto{
		Id:   uid,
		Name: form.Name,
	})
}
