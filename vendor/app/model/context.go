package model

import "database/sql"

type Context struct {
	SQL *sql.DB
}
