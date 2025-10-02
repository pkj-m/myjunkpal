package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type JSONStore struct {
	mu       sync.RWMutex
	dataPath string
}

func NewJSONStore(dataPath string) *JSONStore {
	return &JSONStore{
		dataPath: dataPath,
	}
}

func (s *JSONStore) LoadFromFile(filename string, v interface{}) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	filepath := fmt.Sprintf("%s/%s", s.dataPath, filename)

	data, err := os.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			// Return empty array if file doesn't exist
			return nil
		}
		return err
	}

	if len(data) == 0 {
		return nil
	}

	return json.Unmarshal(data, v)
}

func (s *JSONStore) SaveToFile(filename string, v interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	filepath := fmt.Sprintf("%s/%s", s.dataPath, filename)

	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, data, 0644)
}

func (s *JSONStore) EnsureDataDir() error {
	return os.MkdirAll(s.dataPath, 0755)
}
