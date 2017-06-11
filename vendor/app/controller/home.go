package controller

import (
	"app/session"
	"database/sql"
	"net/http"
)

func (c *Context) HomeGet(w http.ResponseWriter, r *http.Request) {
	s, err := session.DefaultSessionManager().StartSession(w, r)
	if err != nil {
		InternalServerError(w, r)
		return
	}
	uidf, ok := s.Values["user_id"].(float64)
	uid := int64(uidf)
	if !ok {
		c.LandingGet(w, r)
		return
	}

	orgs, err := c.m.OrgsByUserId(uid)
	if err == sql.ErrNoRows {
		c.LandingGet(w, r)
		return
	} else if len(orgs) == 1 {
		http.Redirect(w, r, "/"+orgs[0].Key, http.StatusFound)
	}

	err = c.r.Render(w, "home.html", nil)
	if err != nil {
		InternalServerError(w, r)
	}
}

func (c *Context) LandingGet(w http.ResponseWriter, r *http.Request) {
	err := c.r.Render(w, "landing.html", nil)
	if err != nil {
		InternalServerError(w, r)
	}
}
