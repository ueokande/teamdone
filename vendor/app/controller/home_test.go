package controller

import (
	"app/shared"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func cookieByName(r *http.Response, name string) *http.Cookie {
	for _, c := range r.Cookies() {
		if c.Name == name {
			return c
		}
	}
	return nil
}

func testLogin(r *http.Request, uid int64) error {
	s, err := context.s.StartSession(&httptest.ResponseRecorder{}, r)
	if err != nil {
		return err
	}
	s.Values["user_id"] = uid
	err = context.s.Storage.SessionUpdate(s)
	if err != nil {
		return err
	}
	r.AddCookie(&http.Cookie{
		Name:   "session",
		Value:  s.Id,
		MaxAge: 3600,
	})
	return nil
}

func TestHomeGet_NewUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	context.HomeGet(rec, req)
	resp := rec.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Unexpected status code:", resp.StatusCode)
	}

	c := cookieByName(resp, "session")
	if c == nil {
		t.Fatal("Unexpected response header", resp.Header)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if string(body) != "landing.html" {
		t.Fatal("Unexpected response body:", body)
	}
}

func TestHomeGet_OneOrgUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	uid, err := context.m.UserCreate("alice")
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	err = testLogin(req, uid)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	key := shared.RandomKey()
	oid, err := context.m.OrgCreate("wanderland", key)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	err = context.m.MemberCreate(oid, uid)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	context.HomeGet(rec, req)
	resp := rec.Result()

	if resp.StatusCode != http.StatusFound && resp.Header.Get("Location") != "/"+key {
		t.Fatal("Unexpected status code:", resp.StatusCode)
	}
}

func TestHomeGet_TwoOrgUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	uid, err := context.m.UserCreate("alice")
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	err = testLogin(req, uid)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	for _, name := range []string{"wonderland", "glass"} {
		oid, err := context.m.OrgCreate(name, shared.RandomKey())
		if err != nil {
			t.Fatal("Unexpected error:", err)
		}

		err = context.m.MemberCreate(oid, uid)
		if err != nil {
			t.Fatal("Unexpected error:", err)
		}
	}

	context.HomeGet(rec, req)
	resp := rec.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Unexpected status code:", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if string(body) != "home.html" {
		t.Fatal("Unexpected response body:", body)
	}
}
