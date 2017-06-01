package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOrgGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	OrgGet("my-key", rec, req)

	if code := rec.Code; code != http.StatusOK {
		t.Fatal("Unexpected staus code:", code)
	}
	if body := string(rec.Body.Bytes()); body != "org.html" {
		t.Fatal("Unexpected body:", body)
	}
}
