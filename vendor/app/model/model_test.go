package model

import (
	"app/shared/database"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
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
