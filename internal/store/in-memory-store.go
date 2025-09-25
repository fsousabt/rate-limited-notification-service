package store

import "sync"

type InMemoryStore struct {
	data map[string]Bucket
	mu   sync.Mutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{data: make(map[string]Bucket)}
}

func (s *InMemoryStore) Get(key string) (Bucket, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	bucket, ok := s.data[key]
	return bucket, ok

}
func (s *InMemoryStore) Set(key string, item Bucket) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = item
	return true
}
