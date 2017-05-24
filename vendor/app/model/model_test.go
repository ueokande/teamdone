package model

import (
	"app/shared/database"
	"fmt"
	"os"
	"testing"
)

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

	os.Exit(m.Run())
}
