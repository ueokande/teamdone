package session

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

var ErrNoSession = errors.New("no session")

type Session struct {
	Id     string
	Values map[string]interface{}
}

type SessionStorage interface {
	SessionInit(sid string) (*Session, error)
	SessionRead(sid string) (*Session, error)
	SessionUpdate(*Session) error
	SessionDestroy(sid string) error
}

type Manager struct {
	CookieName string
	Storage    SessionStorage
	LifeTime   time.Duration
}

var defaultSessionManager = &Manager{
	CookieName: "session",
	Storage:    NewMembachedSessionStorage("localhost:11211", 30*24*time.Hour),
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

func newCookie(name string, id string, age time.Duration) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(id),
		Path:     "/",
		HttpOnly: true,
		MaxAge:   int(age),
	}
}

func (m *Manager) newSession(w http.ResponseWriter) (*Session, error) {
	sid := generateId()
	s, err := m.Storage.SessionInit(sid)
	if err != nil {
		return nil, err
	}
	cookie := newCookie(m.CookieName, sid, m.LifeTime)
	http.SetCookie(w, cookie)
	return s, nil
}

func expiredCookie(name string) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now(),
		MaxAge:   -1,
	}
}

func (m *Manager) StartSession(w http.ResponseWriter, r *http.Request) (*Session, error) {
	cookie, err := r.Cookie(m.CookieName)
	if err == http.ErrNoCookie || cookie.Value == "" {
		return m.newSession(w)
	} else if err != nil {
		return nil, err
	}

	sid, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		return nil, err
	}
	s, err := m.Storage.SessionRead(sid)
	if err == ErrNoSession {
		return m.newSession(w)
	} else if err != nil {
		return nil, err
	}
	http.SetCookie(w, newCookie(m.CookieName, cookie.Value, m.LifeTime))
	return s, nil
}

func (m *Manager) DestroySession(w http.ResponseWriter, r *http.Request) error {
	http.SetCookie(w, expiredCookie(m.CookieName))
	cookie, err := r.Cookie(m.CookieName)
	if err == http.ErrNoCookie {
		return nil
	} else if err != nil {
		return err
	}

	err = m.Storage.SessionDestroy(cookie.Value)
	if err == ErrNoSession {
		return nil
	}
	return err
}
