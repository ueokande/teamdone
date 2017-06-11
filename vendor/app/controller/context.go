package controller

import (
	"app/model"
	"database/sql"
)

type Context struct {
	m *model.Context
}

func NewContext(db *sql.DB) *Context {
	return &Context{
		m: &model.Context{
			SQL: db,
		},
	}
}
