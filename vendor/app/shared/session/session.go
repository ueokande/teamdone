package session

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/labstack/echo"
)

var ErrNoSuchValue = errors.New("no such value")
var ErrNoSession = errors.New("no session")

type Manager struct {
	CookieName string
	Storage    SessionStorage
	LifeTime   time.Duration
}

var defaultSessionManager = &Manager{
	CookieName: "session",
	Storage:    NewMemorySessionStorage(),
	LifeTime:   30 * 24 * time.Hour,
}

func generateId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func DefaultSessionManager() *Manager {
	return defaultSessionManager
}

func (m *Manager) newSession(c echo.Context) (Session, error) {
	sid := generateId()
	sess, err := m.Storage.SessionInit(sid)
	if err != nil {
		return nil, err
	}
	c.SetCookie(&http.Cookie{
		Name:     m.CookieName,
		Value:    url.QueryEscape(sid),
		Path:     "/",
		HttpOnly: true,
		MaxAge:   int(m.LifeTime),
	})
	return sess, nil
}

func (m *Manager) StartSession(c echo.Context) (Session, error) {
	cookie, err := c.Cookie(m.CookieName)
	if err == http.ErrNoCookie || err == echo.ErrCookieNotFound || cookie.Value == "" {
		return m.newSession(c)
	} else if err != nil {
		return nil, err
	}

	sid, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		return nil, err
	}
	sess, err := m.Storage.SessionRead(sid)
	if err == ErrNoSession {
		return m.newSession(c)
	} else if err != nil {
		return nil, err
	}
	return sess, nil
}

func (m *Manager) SessionDestroy(c echo.Context) error {
	cookie, err := c.Cookie(m.CookieName)
	if err == echo.ErrCookieNotFound {
		return nil
	} else if err != nil {
		return err
	}
	return m.Storage.SessionDestroy(cookie.Value)
}

func (m *Manager) GC() {
	m.Storage.SessionGC(m.LifeTime)
}
