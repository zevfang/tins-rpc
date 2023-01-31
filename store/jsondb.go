package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
)

type base struct {
	table    map[string]interface{}
	filePath string
	lock     sync.RWMutex
}

func checkErr(errMsg error) {
	if errMsg != nil {
		fmt.Println("run error:", errMsg)
		panic(errMsg)
	}
}

func isExist(fileName string) bool {
	f, err := os.Stat(fileName)
	if err == nil {
		if !f.IsDir() {
			return true
		}
	}
	return false
}

func (db *base) initialize(fileName string) {
	if path.Ext(fileName) != ".json" {
		db.filePath = fileName + ".json"
	} else {
		db.filePath = fileName
	}
	exist := isExist(db.filePath)
	if !exist {
		f, err := os.Create(db.filePath)
		defer f.Close()
		checkErr(err)
	}
	db.syncData()
}

func (db *base) syncData() {
	f, err := os.OpenFile(db.filePath, os.O_RDONLY, 0600)
	defer f.Close()
	checkErr(err)
	contentByte, err := ioutil.ReadAll(f)
	checkErr(err)
	if len(contentByte) != 0 {
		err = json.Unmarshal(contentByte, &db.table)
		checkErr(err)
	} else {
		db.table = make(map[string]interface{})
	}
}

func (db *base) Save() *base {
	f, err := os.OpenFile(db.filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	checkErr(err)
	data, err := json.Marshal(db.table)
	checkErr(err)
	_, err = f.Write(data)
	checkErr(err)
	return db
}

func (db *base) Write(key string, value interface{}) *base {
	db.lock.Lock()
	db.table[key] = value
	db.lock.Unlock()
	return db
}

func (db *base) Read(key string) interface{} {
	return db.table[key]
}

func (db *base) ReadAll() map[string]interface{} {
	return db.table
}

func (db *base) Del(key string) *base {
	delete(db.table, key)
	return db
}

func (db *base) Clear() *base {
	db.table = make(map[string]interface{})
	return db
}

func NewJsonDB(fileName string) *base {
	db := base{
		table:    make(map[string]interface{}),
		filePath: "",
		lock:     sync.RWMutex{},
	}
	db.initialize(fileName)
	return &db
}
