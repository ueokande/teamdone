package controller

import (
	"app/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSessionGetApi(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	uid, err := model.UserCreate("alice")
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	err = testLogin(req, uid)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	SessionGetApi(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatal("Unexpected status:", rec.Code)
	}
	var u SessionGetApiDto
	err = json.Unmarshal(rec.Body.Bytes(), &u)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if u.Id != uid || u.Name != "alice" {
		t.Fatal("Unexpected user:", u)
	}
}

func TestSessionGetApi_NoSessionUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	SessionGetApi(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatal("Unexpected status:", rec.Code)
	}
}
