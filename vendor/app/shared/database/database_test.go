package database

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	db, err := LoadConfig("./testdata/mysql.json")
	if err != nil {
		t.Fatal("Failed to load config:", err)
	}
	mysql, ok := db.(*MySQL)
	if !ok {
		t.Fatal("Failed to load MySQL configuration")
	}

	expected := MySQL{
		Name:     "teamdone",
		Username: "alice",
		Password: "secret",
		Host:     "localhost",
		Port:     3306,
	}
	if *mysql != expected {
		t.Fatal("Unexpected Configure: ", mysql)
	}
}
