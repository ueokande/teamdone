package controller

import (
	"app/model"
	"app/render"
	"database/sql"
)

type Context struct {
	m *model.Context
	r render.Renderer
}

func NewContext(db *sql.DB, r render.Renderer) *Context {
	return &Context{
		m: &model.Context{SQL: db},
		r: r,
	}
}
