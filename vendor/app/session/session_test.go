package session

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo"
)

func setupSessionTest() error {
	storage := NewMembachedSessionStorage("localhost:11211", time.Minute).(*MemcacheSessionStorage)
	err := storage.set(&Session{
		Id: "current-session",
		Values: map[string]interface{}{
			"user": "alice",
		},
	})
	if err != nil {
		return err
	}
	defaultSessionManager.Storage = storage
	return nil
}

func sessionCookie(resp *http.Response) *http.Cookie {
	for _, cookie := range resp.Cookies() {
		if cookie.Name == defaultSessionManager.CookieName {
			return cookie
		}
	}
	return nil
}

func TestStartSession_StoredSession(t *testing.T) {
	err := setupSessionTest()
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("Cookie", defaultSessionManager.CookieName+"=current-session")
	resp := httptest.ResponseRecorder{}

	s, err := DefaultSessionManager().StartSession(&resp, req)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if s.Id != "current-session" {
		t.Fatal("unexpected session id: ", s.Id)
	}
	if u := s.Values["user"]; u != "alice" {
		t.Fatal("unexpected user: ", u)
	}
	if cookie := sessionCookie(resp.Result()); cookie.Value != "current-session" {
		t.Error("unexpected session cookie: ", cookie)
	}
}

func TestStartSession_NoneCookie(t *testing.T) {
	e := echo.New()
	defer e.Close()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.ResponseRecorder{}

	s, err := DefaultSessionManager().StartSession(&resp, req)
	if err != nil {
		t.Fatal(err)
	}
	cookie := sessionCookie(resp.Result())
	if len(cookie.Value) <= 0 || len(s.Id) <= 0 {
		t.Error("unexpected cookie: ", cookie)
	}
}

func TestStartSession_EmptyCookie(t *testing.T) {
	e := echo.New()
	defer e.Close()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("Cookie", defaultSessionManager.CookieName+"=")
	resp := httptest.ResponseRecorder{}

	s, err := DefaultSessionManager().StartSession(&resp, req)
	if err != nil {
		t.Fatal(err)
	}
	cookie := sessionCookie(resp.Result())
	if len(cookie.Value) <= 0 || len(s.Id) <= 0 {
		t.Error("unexpected cookie: ", cookie)
	}
}

func TestStartSession_UnknownSession(t *testing.T) {
	e := echo.New()
	defer e.Close()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("Cookie", defaultSessionManager.CookieName+"=gone-session")
	resp := httptest.ResponseRecorder{}

	s, err := DefaultSessionManager().StartSession(&resp, req)
	if err != nil {
		t.Fatal(err)
	}
	if s.Id == "gone-session-id" {
		t.Error("unexpected session id: ", s.Id)
	}
	if cookie := sessionCookie(resp.Result()); cookie.Value == "gone-session" {
		t.Error("unexpected session name: ", cookie.Name)
	}
}

func TestDestroySession(t *testing.T) {
	e := echo.New()
	defer e.Close()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("Cookie", defaultSessionManager.CookieName+"=current-session")
	resp := httptest.ResponseRecorder{}
	err := DefaultSessionManager().DestroySession(&resp, req)
	if err != nil {
		t.Fatal(err)
	}
	if cookie := sessionCookie(resp.Result()); cookie.Name == "gone-session" {
		t.Error("unexpected session name: ", cookie.Name)
	}
}
