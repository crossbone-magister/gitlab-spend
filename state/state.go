package state

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type State struct {
	registered map[string]bool
	path       string
}

func Load(path string) (*State, error) {
	s := &State{registered: make(map[string]bool), path: path}
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return s, nil
	}
	if err != nil {
		return nil, err
	}
	return s, json.Unmarshal(data, &s.registered)
}

func (s *State) IsRegistered(fingerprint string) bool {
	return s.registered[fingerprint]
}

func (s *State) MarkRegistered(fingerprint string) error {
	s.registered[fingerprint] = true
	if err := os.MkdirAll(filepath.Dir(s.path), 0755); err != nil {
		return err
	}
	data, err := json.Marshal(s.registered)
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, data, 0644)
}
