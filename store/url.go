package store

import (
	"encoding/json"
	"fmt"
)

const urlFileName = "url.json"

type UrlStore struct {
	Path string
}

type UrlStoreModel struct {
	Url   string `json:"url"`
	Frame string `json:"frame"`
}

func NewUrlStore() *UrlStore {
	dir := getLocalDir()
	s := &UrlStore{}
	s.Path = fmt.Sprintf("%s%s", dir, urlFileName)
	return s
}

func (s *UrlStore) Get(key string) UrlStoreModel {
	var res UrlStoreModel
	val := NewJsonDB(s.Path).Read(key)
	if val == nil {
		return res
	}
	valStr, err := json.Marshal(val)
	if err != nil {
		return res
	}
	_ = json.Unmarshal(valStr, &res)
	return res
}

func (s *UrlStore) Set(key string, value UrlStoreModel) {
	NewJsonDB(s.Path).Write(key, value).Save()
}

func (s *UrlStore) Delete(key string) {
	NewJsonDB(s.Path).Del(key).Save()
}

func (s *UrlStore) Clear() {
	NewJsonDB(s.Path).Clear().Save()
}
