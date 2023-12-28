package nopanic

import "sync"

type Slice[T any] struct {
	values []T
	mux    sync.RWMutex
}

func (s *Slice[T]) Add(vs ...T) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.values = append(s.values, vs...)
}

func (s *Slice[T]) All() []T {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.values
}

func (s *Slice[T]) Purge() {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.values = nil
}
