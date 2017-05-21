package session

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/labstack/echo"
)

func TestStartSession_StoredSession(t *testing.T) {
	e := echo.New()
	defer e.Close()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add(echo.HeaderCookie, defaultSessionManager.CookieName+"=my-session-id")
	resp := httptest.ResponseRecorder{}

	c := e.NewContext(req, &resp)
	defaultSessionManager.Storage.(*MemorySessionStorage).sessions["my-session-id"] = &MemorySession{
		sid: "my-session-id",
		values: map[interface{}]interface{}{
			"user": "alice",
		},
		mutex: new(sync.Mutex),
	}
	sess, err := DefaultSessionManager().StartSession(c)
	if err != nil {
		t.Fatal(err)
	}
	if sess.SessionID() != "my-session-id" {
		t.Error("unexpected session id: ", sess.SessionID())
	}
	if user, err := sess.Get("user"); err != nil {
		t.Fatal(err)
	} else if user != "alice" {
		t.Error("unexpected user: ", user)
	}
}

func TestStartSession_UnknownSession(t *testing.T) {
	e := echo.New()
	defer e.Close()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add(echo.HeaderCookie, defaultSessionManager.CookieName+"=old-session-id")
	resp := httptest.ResponseRecorder{}

	c := e.NewContext(req, &resp)
	sess, err := DefaultSessionManager().StartSession(c)
	if err != nil {
		t.Fatal(err)
	}
	if sess.SessionID() == "old-session-id" {
		t.Error("unexpected session id: ", sess.SessionID())
	}
}
