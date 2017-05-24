package model

import (
	"testing"
)

func TestUserCreate(t *testing.T) {
	id, err := UserCreate("alice")
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	u, err := UserById(id)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	if u.Name != "alice" {
		t.Fatal("Unexpected user name:", u.Name)

	}
}
