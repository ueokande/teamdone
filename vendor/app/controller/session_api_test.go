package controller

import (
	"app/model"
	"bytes"
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
	if u.UserId != uid || u.UserName != "alice" {
		t.Fatal("Unexpected user:", u)
	}
}

func TestSessionGetApi_NoSessionUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	SessionGetApi(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatal("Unexpected status:", rec.Code)
	}
	if body := string(rec.Body.Bytes()); body != `{}` {
		t.Fatal("Unexpected body:", body)
	}
}

func TestSessionCreateApiTest(t *testing.T) {
	cases := []struct {
		req  string
		code int
	}{
		{`{}`, http.StatusBadRequest},
		{`{ "Name" : "" }`, http.StatusBadRequest},
		{`{ "Name"}`, http.StatusBadRequest},
		{`{ "Name": "alice" }`, http.StatusOK},
	}

	for _, c := range cases {
		req := httptest.NewRequest(http.MethodGet, "/i/session/create", bytes.NewReader([]byte(c.req)))
		rec := httptest.NewRecorder()

		SessionCreateApi(rec, req)

		if rec.Code != http.StatusOK {
			continue
		}
		resp := make(map[string]interface{})
		err := json.NewDecoder(bytes.NewReader(rec.Body.Bytes())).Decode(&resp)
		if err != nil {
			t.Fatal("Unexpected error:", err)
		}
		if _, ok := resp["Id"]; !ok {
			t.Fatal("Unexpected response:", resp)
		}
		if _, ok := resp["Name"]; !ok {
			t.Fatal("Unexpected response:", resp)
		}
	}

}
