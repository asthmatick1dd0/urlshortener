package repository

import (
	"fmt"
	"sync"
)

type URLRepository interface {
	Save(code, url string) error
	Get(code string) (string, error)
}

type InMemoryURLRepository struct {
	mu    sync.RWMutex
	store map[string]string
}

func NewInMemoryURLRepository() *InMemoryURLRepository {
	return &InMemoryURLRepository{
		store: make(map[string]string),
	}
}

func (r *InMemoryURLRepository) Save(code, url string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.store[code] = url
	return nil
}

func (r *InMemoryURLRepository) Get(code string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	url, exists := r.store[code]
	if !exists {
		return "", fmt.Errorf("URL not found")
	}
	return url, nil
}
