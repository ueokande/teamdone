package model

import (
	"app/shared/database"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

var context *Context

func initializeDB() error {
	conf, err := database.LoadConfig("../../../config/test.json")
	if err != nil {
		return err
	}
	db, err := database.Connect(conf)
	if err != nil {
		return err
	}
	context = &Context{SQL: db}
	return nil
}

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())

	err := initializeDB()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}
