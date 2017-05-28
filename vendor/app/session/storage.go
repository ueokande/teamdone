package session

import (
	"encoding/json"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type MemcacheSessionStorage struct {
	client   *memcache.Client
	lifetime time.Duration
}

func NewMembachedSessionStorage(server string, lifetime time.Duration) SessionStorage {
	return &MemcacheSessionStorage{
		client:   memcache.New(server),
		lifetime: lifetime,
	}
}

func (p *MemcacheSessionStorage) get(sid string) (*Session, error) {
	item, err := p.client.Get(sid)
	if err != nil {
		return nil, err
	}
	err = p.client.Touch(sid, int32(p.lifetime/1e9))
	if err != nil {
		return nil, err
	}
	s := Session{
		Id: sid,
	}
	err = json.Unmarshal(item.Value, &s.Values)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
func (p *MemcacheSessionStorage) set(s *Session) error {
	data, err := json.Marshal(s.Values)
	if err != nil {
		return err
	}
	item := &memcache.Item{
		Key:        s.Id,
		Value:      data,
		Expiration: int32(p.lifetime / 1e9),
	}
	return p.client.Set(item)
}

func (p *MemcacheSessionStorage) delete(sid string) error {
	return p.client.Delete(sid)
}

func (p *MemcacheSessionStorage) SessionInit(sid string) (*Session, error) {
	s := &Session{
		Id:     sid,
		Values: make(map[string]interface{}),
	}
	err := p.set(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (p *MemcacheSessionStorage) SessionRead(sid string) (*Session, error) {
	s, err := p.get(sid)
	if err == memcache.ErrCacheMiss {
		return nil, ErrNoSession
	}
	return s, err
}

func (p *MemcacheSessionStorage) SessionUpdate(s *Session) error {
	return p.set(s)
}

func (p *MemcacheSessionStorage) SessionDestroy(sid string) error {
	err := p.delete(sid)
	if err == memcache.ErrCacheMiss {
		return ErrNoSession
	}
	return err
}
