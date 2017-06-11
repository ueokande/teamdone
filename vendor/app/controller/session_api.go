package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type SessionGetApiDto struct {
	UserId   int64
	UserName string
}

type SessionCreateApiForm struct {
	UserName string
}

type SessionCreateApiDto struct {
	UserId   int64
	UserName string
}

func (c *Context) SessionGetApi(w http.ResponseWriter, r *http.Request) {
	s, err := c.s.StartSession(w, r)
	if err != nil {
		InternalServerError(w, r)
		return
	}
	uidf, ok := s.Values["user_id"].(float64)
	uid := int64(uidf)
	if !ok {
		jsonOk(w, struct{}{})
		return
	}
	u, err := c.m.UserById(uid)
	if err == sql.ErrNoRows {
		jsonOk(w, struct{}{})
		return
	} else if err != nil {
		InternalServerError(w, r)
		return
	}

	jsonOk(w, SessionGetApiDto{
		UserId:   u.Id,
		UserName: u.Name,
	})
}

func (c *Context) SessionCreateApi(w http.ResponseWriter, r *http.Request) {
	s, err := c.s.StartSession(w, r)
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
	if len(form.UserName) == 0 {
		jsonError(w, "user name is required", http.StatusBadRequest)
		return
	}

	uid, err = c.m.UserCreate(form.UserName)
	if err != nil {
		InternalServerError(w, r)
		return
	}

	s.Values["user_id"] = uid
	err = c.s.Storage.SessionUpdate(s)
	if err != nil {
		InternalServerError(w, r)
		return
	}

	jsonOk(w, SessionCreateApiDto{
		UserId:   uid,
		UserName: form.UserName,
	})
}
