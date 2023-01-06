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
	// GUI hosts
	Uris []Uri `json:"uris"`
}

type Uri struct {
	Proto string `json:"proto"`
	Uri   string `json:"uri"`
}

func NewStorage() *Storage {
	storageDir := getStorageDir()
	s := &Storage{}
	s.FilePath = fmt.Sprintf("%s%s", storageDir, storageFile)
	exists, _ := HasDir(s.FilePath)
	if !exists {
		CreateDir(storageDir)
		data := LocalData{
			Uris: []Uri{},
		}
		b, _ := json.Marshal(&data)
		SaveFile(s.FilePath, b)
	}
	return s
}

func (s *Storage) GetUris(proto string) Uri {
	data := s.getData()
	for _, uri := range data.Uris {
		if proto == uri.Proto {
			return uri
		}
	}
	return Uri{}
}

func (s *Storage) SetUris(uri Uri) {
	data := s.getData()
	u := s.GetUris(uri.Proto)
	if len(u.Proto) > 0 {
		return
	}
	data.Uris = append(data.Uris, uri)
	b, _ := json.Marshal(&data)
	SaveFile(s.FilePath, b)
}

func (s *Storage) getData() LocalData {
	b := ReadFile(s.FilePath)
	data := LocalData{}
	_ = json.Unmarshal(b, &data)
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
