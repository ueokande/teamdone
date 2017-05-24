package session

import (
	"sync"
	"testing"
	"time"
)

func TestSessionInit(t *testing.T) {
	s := NewMemorySessionStorage()
	session, err := s.SessionInit("my-session-id")

	if err != nil {
		t.Fatal(err)
	}
	if session.SessionID() != "my-session-id" {
		t.Error("Unexpected session ID")
	}
}

func TestSessionRead(t *testing.T) {
	s := &MemorySessionStorage{
		lock: new(sync.Mutex),
		sessions: map[string]*MemorySession{
			"my-session-id": &MemorySession{
				sid: "my-session-id",
			},
		},
	}

	cases := []struct {
		sid string
		err error
	}{
		{"my-session-id", nil},
		{"my-gone-session-id", ErrNoSession},
	}

	for _, c := range cases {
		sess, err := s.SessionRead(c.sid)
		if err != c.err {
			t.Error("unexpected error: ", err)
		}
		if err == nil && sess.SessionID() != c.sid {
			t.Error("unexpected session ID: ", sess.SessionID())
		}
	}
}

func TestSessionDestroy(t *testing.T) {
	s := &MemorySessionStorage{
		lock: new(sync.Mutex),
		sessions: map[string]*MemorySession{
			"my-session-id": &MemorySession{
				sid: "my-session-id",
			},
		},
	}

	cases := []struct {
		sid string
	}{
		{"my-session-id"},
		{"my-gone-session-id"},
	}

	for _, c := range cases {
		err := s.SessionDestroy(c.sid)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := s.sessions[c.sid]; ok {
			t.Error("unexpected session contained")
		}
	}
}

func TestSessionGC(t *testing.T) {
	s := &MemorySessionStorage{
		lock:     new(sync.Mutex),
		sessions: make(map[string]*MemorySession),
	}
	cases := []struct {
		sid          string
		lastAccessed time.Time
		exists       bool
	}{
		{"old-session", time.Now().Add(-60 * time.Minute), false},
		{"new-session", time.Now().Add(-10 * time.Minute), true},
	}

	for _, c := range cases {
		s.sessions[c.sid] = &MemorySession{
			sid:          c.sid,
			lastAccessed: c.lastAccessed,
		}
	}
	s.SessionGC(30 * time.Minute)

	for _, c := range cases {
		_, exists := s.sessions[c.sid]
		if exists != c.exists {
			t.Error("unexpected sesson: ", c.sid)
		}
	}

}

func newSession() *MemorySession {
	return &MemorySession{
		sid: "my-session-id",
		values: map[interface{}]interface{}{
			"email": "abc@example.com",
		},
		lastAccessed: time.Now().Add(-30 * time.Minute),
		mutex:        new(sync.Mutex),
	}
}

func TestGet(t *testing.T) {
	s := newSession()
	cases := []struct {
		key   string
		err   error
		value string
	}{
		{"email", nil, "abc@example.com"},
		{"userid", ErrNoSuchValue, ""},
	}

	for _, c := range cases {
		value, err := s.Get(c.key)
		if err != c.err {
			t.Error("unexpected error: ", err)
		}
		if err == nil && value != c.value {
			t.Error("unexpected value: ", value)
		}
	}
}

func TestSet(t *testing.T) {
	s := newSession()
	cases := []struct {
		key   string
		value string
	}{
		{"email", "abc@example.com"},
		{"2ndemail", "new@example.com"},
	}

	for _, c := range cases {
		err := s.Set(c.key, c.value)
		if err != nil {
			t.Fatal(err)
		}
		if s.values[c.key] != c.value {
			t.Error("Unexpected value: ", s.values[c.key])
		}
	}
}

func TestDelete(t *testing.T) {
	s := newSession()
	cases := []struct {
		key string
	}{
		{"email"}, {"2ndemail"},
	}

	for _, c := range cases {
		err := s.Delete(c.key)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := s.values[c.key]; ok {
			t.Error("Unexpected key: ", c.key)
		}
	}
}
