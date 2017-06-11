package controller

import (
	"app/model"
	"app/render"
	"app/session"
	"database/sql"
)

type Context struct {
	m *model.Context
	s *session.Manager
	r render.Renderer
}

func NewContext(db *sql.DB, s *session.Manager, r render.Renderer) *Context {
	return &Context{
		m: &model.Context{SQL: db},
		s: s,
		r: r,
	}
}
