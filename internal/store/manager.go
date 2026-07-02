package store

import (
	"sync"
)

// StoreManager maintains multiple Store instances for different project directories.
// This allows a single global MCP server instance to service multiple novels simultaneously.
type StoreManager struct {
	defaultDir string
	stores     map[string]*Store
	mu         sync.Mutex
}

// NewStoreManager initializes a new StoreManager with a fallback default directory.
func NewStoreManager(defaultDir string) *StoreManager {
	return &StoreManager{
		defaultDir: defaultDir,
		stores:     make(map[string]*Store),
	}
}

// GetStore retrieves an existing Store for the given directory or creates a new one.
// If dir is empty, it falls back to the defaultDir configured at startup.
func (m *StoreManager) GetStore(dir string) *Store {
	if dir == "" {
		dir = m.defaultDir
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if s, ok := m.stores[dir]; ok {
		return s
	}

	s := NewStore(dir)
	m.stores[dir] = s
	return s
}
