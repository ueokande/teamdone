package controller

import (
	"app/model"
	"app/session"
	"database/sql"
	"net/http"
)

type SessionGetApiDto struct {
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
