package controller

import (
	"app/model"
	"app/session"
	"app/shared/database"
	"database/sql"
	"fmt"
	"io"
	"math/rand"
	"os"
	"testing"
	"time"
)

type MockRender struct{}

var context *Context

func (t *MockRender) Render(w io.Writer, name string, data interface{}) error {
	fmt.Fprintf(w, name)
	return nil
}

func initializeDB() (*sql.DB, error) {
	conf, err := database.LoadConfig("../../../config/test.json")
	if err != nil {
		return nil, err
	}
	return database.Connect(conf)
}

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())

	db, err := initializeDB()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	context = &Context{
		m: &model.Context{SQL: db},
		s: &session.Manager{
			CookieName: "session",
			Storage:    session.NewMembachedSessionStorage("localhost:11211", 30*24*time.Hour),
			LifeTime:   30 * 24 * time.Hour,
		},
		r: &MockRender{},
	}

	os.Exit(m.Run())
}
