package controller

import (
	"app/model"
	"app/render"
	"app/session"
	"database/sql"
	"net/http"
)

func HomeGet(w http.ResponseWriter, r *http.Request) {
	s, err := session.DefaultSessionManager().StartSession(w, r)
	if err != nil {
		InternalServerError(w, r)
		return
	}
	uid, ok := s.Values["user_id"].(int64)
	if !ok {
		LandingGet(w, r)
		return
	}

	orgs, err := model.OrgsByUserId(uid)
	if err != sql.ErrNoRows {
		LandingGet(w, r)
		return
	} else if len(orgs) == 1 {
		http.Redirect(w, r, "/"+orgs[0].Key, http.StatusFound)
	}

	err = render.Render(w, "home.html", nil)
	if err != nil {
		InternalServerError(w, r)
	}
}

func LandingGet(w http.ResponseWriter, r *http.Request) {
	err := render.Render(w, "landing.html", nil)
	if err != nil {
		InternalServerError(w, r)
	}
}
