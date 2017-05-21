package session

import (
	"sync"
	"time"
)

type Session interface {
	SessionID() string
	Get(key interface{}) (interface{}, error)
	Set(key, value interface{}) error
	Delete(key interface{}) error
}

type SessionStorage interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(lifetime time.Duration)
}

type MemorySessionStorage struct {
	lock     *sync.Mutex
	sessions map[string]*MemorySession
}

func NewMemorySessionStorage() SessionStorage {
	return &MemorySessionStorage{
		lock:     new(sync.Mutex),
		sessions: make(map[string]*MemorySession),
	}
}

func (p *MemorySessionStorage) SessionInit(sid string) (Session, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	sess := &MemorySession{
		sid:          sid,
		lastAccessed: time.Now(),
		values:       make(map[interface{}]interface{}, 0),
		mutex:        new(sync.Mutex),
	}
	p.sessions[sid] = sess
	return sess, nil
}

func (p *MemorySessionStorage) SessionRead(sid string) (Session, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if sess, ok := p.sessions[sid]; ok {
		return sess, nil
	}
	return nil, ErrNoSession
}

func (p *MemorySessionStorage) SessionDestroy(sid string) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	delete(p.sessions, sid)
	return nil
}

func (p *MemorySessionStorage) SessionGC(lifetime time.Duration) {
	now := time.Now()
	for k, v := range p.sessions {
		if v.lastAccessed.Add(lifetime).Before(now) {
			p.SessionDestroy(k)
		}
	}
}

type MemorySession struct {
	sid          string
	values       map[interface{}]interface{}
	lastAccessed time.Time
	mutex        *sync.Mutex
}

func (s *MemorySession) SessionID() string {
	return s.sid
}

func (s *MemorySession) Get(key interface{}) (interface{}, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.lastAccessed = time.Now()
	value, ok := s.values[key]
	if !ok {
		return nil, ErrNoSuchValue
	}
	return value, nil
}

func (s *MemorySession) Set(key, value interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.lastAccessed = time.Now()
	s.values[key] = value
	return nil
}

func (s *MemorySession) Delete(key interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.lastAccessed = time.Now()
	delete(s.values, key)
	return nil
}
