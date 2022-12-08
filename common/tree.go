package common

import (
	"fmt"
)

type TreeData struct {
	TreeData map[string][]string
	JsonData map[string]string // {Service.Rpc}:{RequestJsonString}
}

func NewTreeData() *TreeData {
	t := &TreeData{
		TreeData: make(map[string][]string),
		JsonData: make(map[string]string),
	}
	t.TreeData[""] = []string{} //初始化根节点
	return t
}

func (t *TreeData) GetProtoData(filePath string) error {
	result, err := GerProtoData(filePath)
	if err != nil {
		return err
	}
	parent := fmt.Sprintf("[S] %s.%s", result.PackageName, result.ServiceName)
	t.Clear(parent)
	t.AppendParentTreeNode(parent)
	for rpcName, reqJson := range result.RequestList {
		rpcName = fmt.Sprintf("[M] %s.%s", result.ServiceName, rpcName)
		t.AppendChildTreeNode(parent, rpcName)
		t.SetRequestJson(rpcName, reqJson)
	}
	return nil
}

func (t *TreeData) Exist(node string) bool {
	if _, ok := t.TreeData[node]; ok {
		return true
	}
	return false
}

func (t *TreeData) Clear(node string) {
	// 检索重复元素
	var deleteIndex int = -1
	for i, s := range t.TreeData[""] {
		if s == node {
			deleteIndex = i
		}
	}
	if deleteIndex > -1 {
		// 清空父元素
		list := t.TreeData[""]
		newList := append(list[:deleteIndex], list[(deleteIndex+1):]...)
		t.TreeData[""] = newList
		// 清空子元素
		if _, ok := t.TreeData[node]; ok {
			t.TreeData[node] = []string{}
		}
	}
}

func (t *TreeData) AppendParentTreeNode(node ...string) {
	t.TreeData[""] = append(t.TreeData[""], node...)
}

func (t *TreeData) AppendChildTreeNode(parent string, node ...string) {
	t.TreeData[parent] = append(t.TreeData[parent], node...)
}

func (t *TreeData) SetRequestJson(rpcName string, reqJson string) {
	t.JsonData[rpcName] = reqJson
}

func (t *TreeData) GetRequestJson(key string) string {
	return t.JsonData[key]
}

func (t *TreeData) RemoveAll() {
	t.TreeData = map[string][]string{"": {}}
	t.JsonData = map[string]string{}
}
