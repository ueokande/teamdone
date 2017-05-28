package session

import (
	"testing"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

func setupStorageTest() (SessionStorage, error) {
	client := memcache.New("localhost:11211")
	err := client.DeleteAll()
	if err != nil {
		return nil, err
	}
	client.Set(&memcache.Item{
		Key:        "current-session",
		Value:      []byte(`{"user_id": 100}`),
		Expiration: int32(60),
	})
	storage := &MemcacheSessionStorage{
		client:   client,
		lifetime: time.Minute,
	}
	return storage, err
}

func TestSessionInit(t *testing.T) {
	storage, err := setupStorageTest()
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	s, err := storage.SessionInit("my-session-id")
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if s.Id != "my-session-id" {
		t.Error("Unexpected session Id:", s.Id)
	}

	s, err = storage.SessionRead("my-session-id")
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if s.Id != "my-session-id" {
		t.Error("Unexpected session Id:", s.Id)
	}
}

func TestSessionRead(t *testing.T) {
	storage, err := setupStorageTest()
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	cases := []struct {
		sid string
		err error
		uid int
	}{
		{"current-session", nil, 100},
		{"gone-session-id", ErrNoSession, 0},
	}

	for _, c := range cases {
		s, err := storage.SessionRead(c.sid)
		if err != c.err {
			t.Error("unexpected error: ", err)
		}
		if err != nil {
			continue
		}
		if s.Id != c.sid {
			t.Error("unexpected session Id: ", s.Id)
		}
		if s.Values["user_id"] != float64(c.uid) {
			t.Error("unexpected user id: ", s.Values["user_id"])
		}
	}
}

func SessionUpdate(t *testing.T) {
	storage, err := setupStorageTest()
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	s, err := storage.SessionRead("current-session")
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	s.Values["user_id"] = 123
	err = storage.SessionUpdate(s)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	s, err = storage.SessionRead("current-session")
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if s.Values["user_id"] != 123 {
		t.Error("Unexpected session Id:", s.Id)
	}
}

func SessionUpdateGoneSession(t *testing.T) {
	storage, err := setupStorageTest()
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	err = storage.SessionUpdate(&Session{
		Id: "gone-session",
	})
	if err != ErrNoSession {
		t.Fatal("Unexpected error:", err)
	}
}

func TestSessionDestroyCurrentSession(t *testing.T) {
	storage, err := setupStorageTest()
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	err = storage.SessionDestroy("current-session")
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	_, err = storage.SessionRead("current-session")
	if err != ErrNoSession {
		t.Fatal("Unexpected error:", err)
	}
}

func TestSessionDestroyGoneSession(t *testing.T) {
	storage, err := setupStorageTest()
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	err = storage.SessionDestroy("gone-session")
	if err != ErrNoSession {
		t.Fatal("Unexpected error:", err)
	}
}
