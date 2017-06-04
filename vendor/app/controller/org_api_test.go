package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestOrgCreateApi(t *testing.T) {

	cases := []struct {
		req  string
		code int
	}{
		{`{}`, http.StatusBadRequest},
		{`{ "OrgName" : "" }`, http.StatusBadRequest},
		{`{ "OrgName" `, http.StatusBadRequest},
		{`{ "OrgName" : "wonderland" }`, http.StatusOK},
	}

	for _, c := range cases {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(c.req)))
		rec := httptest.NewRecorder()

		OrgCreateApi(rec, req)

		if code := rec.Code; code != c.code {
			t.Fatal("Unexpected staus code:", code)
		}
		if rec.Code != http.StatusOK {
			continue
		}
		if body := string(rec.Body.Bytes()); !strings.Contains(body, `"Key"`) {
			t.Fatal("Unexpected body:", body)
		}
	}
}
