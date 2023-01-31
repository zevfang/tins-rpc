package store

import (
	"fmt"
)

const treeFileName = "tree.json"

type TreeStore struct {
	Path string
}

func NewTreeStore() *TreeStore {
	dir := getLocalDir()
	s := &TreeStore{}
	s.Path = fmt.Sprintf("%s%s", dir, treeFileName)
	return s
}

func (s *TreeStore) Get(key string) string {
	val := NewJsonDB(s.Path).Read(key)
	if val == nil {
		return ""
	}
	return val.(string)
}

func (s *TreeStore) Set(key string, value interface{}) {
	NewJsonDB(s.Path).Write(key, value).Save()
}

func (s *TreeStore) Delete(key string) {
	NewJsonDB(s.Path).Del(key).Save()
}

func (s *TreeStore) Clear() {
	NewJsonDB(s.Path).Clear().Save()
}

func (s *TreeStore) List() []string {
	var p []string
	values := NewJsonDB(s.Path).ReadAll()
	if values == nil {
		return p
	}
	for _, path := range values {
		p = append(p, path.(string))
	}
	return p
}
