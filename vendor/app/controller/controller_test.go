package controller

import (
	"app/model"
	"app/render"
	"app/shared/database"
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

func initializeDB() error {
	conf, err := database.LoadConfig("../../../config/test.json")
	if err != nil {
		return err
	}
	db, err := database.Connect(conf)
	if err != nil {
		return err
	}
	context = &Context{
		m: &model.Context{SQL: db},
	}
	return nil
}

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())

	err := initializeDB()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	render.DefaultRenderer = &MockRender{}

	os.Exit(m.Run())
}
