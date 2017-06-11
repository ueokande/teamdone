package model

import (
	"app/shared/database"
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

var context *Context

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

	context = &Context{SQL: db}

	os.Exit(m.Run())
}
