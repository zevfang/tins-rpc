package store

import (
	"fmt"
	"os"
	"runtime"
)

type Store struct {
	Url  *UrlStore
	Tree *TreeStore
}

func InitStore() *Store {
	return &Store{
		Url:  NewUrlStore(),
		Tree: NewTreeStore(),
	}
}

func getLocalDir() string {
	userHomedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	switch runtime.GOOS {
	case "darwin":
		return fmt.Sprintf("%s%s", userHomedir, "/.local/share/TinsRPC/")
	case "linux":
		return fmt.Sprintf("%s%s", userHomedir, "/.local/share/TinsRPC/")
	case "windows":
		return fmt.Sprintf("%s%s", userHomedir, "\\.local\\share\\TinsRPC\\")
	}
	return ""
}
