package controller

import (
	"app/render"
	"app/shared/database"
	"fmt"
	"io"
	"os"
	"testing"
)

type MockRender struct{}

func (t *MockRender) Render(w io.Writer, name string, data interface{}) error {
	fmt.Fprintf(w, name)
	return nil
}

func initializeDB() error {
	db, err := database.LoadConfig("../../../config/test.json")
	if err != nil {
		return err
	}
	err = database.Connect(db)
	if err != nil {
		return err
	}
	_, err = database.SQL.Exec("DELETE FROM member")
	if err != nil {
		return err
	}
	_, err = database.SQL.Exec("DELETE FROM task")
	if err != nil {
		return err
	}
	_, err = database.SQL.Exec("DELETE FROM user")
	if err != nil {
		return err
	}
	_, err = database.SQL.Exec("DELETE FROM org")
	if err != nil {
		return err
	}
	return nil
}

func TestMain(m *testing.M) {
	err := initializeDB()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	render.DefaultRenderer = &MockRender{}

	os.Exit(m.Run())
}
