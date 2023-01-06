package data

import (
	"fmt"
	"io/ioutil"
	"os"
)

func HasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

func CreateDir(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf("create dir err -> %v\n", err)
	}
}

func RemoveDir(path string) error {
	_err := os.RemoveAll(path)
	return _err
}

func ReadFile(fileName string) []byte {
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		contentByte, _ := ioutil.ReadAll(f)
		return contentByte
	}
	return nil
}

func SaveFile(fileName string, data []byte) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write(data)
	}
}

func RemoveFile(fileName string) error {
	_err := os.Remove(fileName)
	return _err
}
