package data

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

const storageFile = "data.json"
const (
	WINDOWS_DIR = "\\.local\\share\\TinsRPC\\"
	MAC_DIR     = "/.local/share/TinsRPC/"
	LINUX_DIR   = "/.local/share/TinsRPC/"
)

type Storage struct {
	FilePath string
}

type LocalData struct {
	// GUI hosts k:proto v:uri
	Uris map[string]string `json:"uris"`
}

func NewStorage() *Storage {
	storageDir := getStorageDir()
	s := &Storage{}
	s.FilePath = fmt.Sprintf("%s%s", storageDir, storageFile)
	exists, _ := HasDir(s.FilePath)
	if !exists {
		CreateDir(storageDir)
		data := LocalData{
			Uris: make(map[string]string),
		}
		b, _ := json.Marshal(&data)
		SaveFile(s.FilePath, b)
	}
	return s
}

func (s *Storage) GetUris(proto string) string {
	data := s.getData()
	if _, found := data.Uris[proto]; found {
		return data.Uris[proto]
	}
	return ""
}

func (s *Storage) SetUris(proto, uri string) {
	data := s.getData()
	if data.Uris == nil {
		return
	}
	data.Uris[proto] = uri
	b, _ := json.Marshal(data)
	SaveFile(s.FilePath, b)
}

func (s *Storage) getData() LocalData {
	b := ReadFile(s.FilePath)
	fmt.Println("file->", string(b))
	data := LocalData{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		// 载入默认值
		data.Uris = make(map[string]string)
		dataJson, _ := json.Marshal(&data)
		_ = os.Truncate(s.FilePath, 0) //清空文件
		SaveFile(s.FilePath, dataJson)
	}

	return data
}

func getStorageDir() string {
	userHomedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	switch runtime.GOOS {
	case "darwin":
		return fmt.Sprintf("%s%s", userHomedir, MAC_DIR)
	case "linux":
		return fmt.Sprintf("%s%s", userHomedir, LINUX_DIR)
	case "windows":
		return fmt.Sprintf("%s%s", userHomedir, WINDOWS_DIR)
	}
	return ""
}
