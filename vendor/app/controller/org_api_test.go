package controller

import (
	"app/shared"
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestOrgGetApi(t *testing.T) {
	key := shared.RandomKey()
	context.m.OrgCreate("name", key)

	cases := []struct {
		req  string
		code int
	}{
		{`{ "OrgKey":"` + key + `"}`, 200},
		{`{ "OrgKey":"unregistered-key"}`, 404},
		{`{}`, 404},
		{`uinvalid json`, 400},
	}

	for _, c := range cases {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(c.req)))
		rec := httptest.NewRecorder()

		context.OrgGetApi(rec, req)
		if code := rec.Code; code != c.code {
			t.Fatal("Unexpected staus code:", code)
		}
		if rec.Code != http.StatusOK {
			continue
		}
		if body := string(rec.Body.Bytes()); !strings.Contains(body, `"OrgId"`) {
			t.Fatal("Unexpected body:", body)
		}
	}
}

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

		context.OrgCreateApi(rec, req)

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
