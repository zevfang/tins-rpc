package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	pp "github.com/emicklei/proto"
	"tins-rpc/call"
	"tins-rpc/common"
	tinsTheme "tins-rpc/theme"
)

func menuTree() *widget.Tree {
	tree := widget.NewTree(
		//childUIDs
		func(uid widget.TreeNodeID) []widget.TreeNodeID {
			return MenuTree.Nodes(uid)
		},
		//isBranch
		func(uid widget.TreeNodeID) bool {
			_, b := MenuTree.ProtoData[uid]
			return b
		},
		//create
		func(b bool) fyne.CanvasObject {
			return container.NewHBox(
				widget.NewIcon(tinsTheme.ResourceMSquareIcon),
				widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{}))
		},
		//update
		func(uid widget.TreeNodeID, b bool, object fyne.CanvasObject) {
			switch MenuTree.NodeType(uid) {
			case ProtoProto:
				object.(*fyne.Container).Objects[0].(*widget.Icon).SetResource(tinsTheme.ResourcePSquareIcon)
			case ProtoService:
				object.(*fyne.Container).Objects[0].(*widget.Icon).SetResource(tinsTheme.ResourceSSquareIcon)
			case ProtoMethod:
				object.(*fyne.Container).Objects[0].(*widget.Icon).SetResource(tinsTheme.ResourceMSquareIcon)
			}
			object.(*fyne.Container).Objects[1].(*widget.Label).SetText(uid)
		},
	)
	tree.OnSelected = func(uid string) {
		_type := MenuTree.NodeType(uid)
		if _type == ProtoProto || _type == ProtoService {
			return
		}
		// 删除 New Tab
		for i, item := range globalWin.tabs.Items {
			if item.Text == tinsTheme.WelComeTabTitle {
				globalWin.tabs.RemoveIndex(i)
			}
		}
		// Tab 获取显示名称
		uidTab := MenuTree.TabName(uid)
		// 检测打开并选中
		if _, ok := TabItemList[uidTab]; ok {
			//设置选中
			globalWin.tabs.Select(TabItemList[uidTab].TabItem)
			return
		}
		// 添加选项卡
		tabItem := AppendTabItemView(uidTab, globalWin.tabs)
		// 设置被选中
		tabItem.SelectTree = uidTab

		// 获取选中方法json
		data := MenuTree.RequestJson(uid)
		fmt.Println(uid, data)
		tabItem.RequestText.Text = data
		tabItem.RequestText.Refresh()
		//设置选中
		globalWin.tabs.Select(tabItem.TabItem)
		//保存TabItem
		TabItemList[uidTab] = tabItem
	}
	tree.OpenAllBranches()
	tree.Refresh()
	return tree
}

const (
	ProtoProto   = "proto"   //文件
	ProtoService = "service" //服务
	ProtoMethod  = "method"  //方法
)

type TreeData struct {
	ProtoData map[string][]TreeNode
	JsonData  map[string]string // {Service.Rpc}:{RequestJsonString}
	ProtoFds  map[string]call.ProtoDescriptor
}

type TreeNode struct {
	NodeID  string
	Type    string
	Data    interface{}
	JsonStr string
}

func NewTreeData() *TreeData {
	t := &TreeData{
		ProtoData: make(map[string][]TreeNode),
		ProtoFds:  make(map[string]call.ProtoDescriptor),
	}
	t.ProtoData[""] = make([]TreeNode, 0) //初始化根节点
	return t
}

func (t *TreeData) Append(filePath string) error {
	// 解析proto到definit
	definitions := common.NewDefinitions()
	_ = definitions.ReadFile(filePath)
	// 转换为tree结构
	treeData := t.Parse(definitions)
	for s, nodes := range treeData {
		t.ProtoData[s] = append(t.ProtoData[s], nodes...)
	}
	return nil
}

func (t *TreeData) Parse(d *common.Definitions) map[string][]TreeNode {
	msgJson := common.NewDecoder(d).DecodeAll()
	data := make(map[string][]TreeNode)
	// "":greeter.proto
	data[""] = append(data[""], TreeNode{
		NodeID:  d.GetFileName(),
		Type:    ProtoProto,
		Data:    nil,
		JsonStr: "",
	})
	for svcName, svcDef := range d.GetServices() {
		//greeter.proto:ct.Ct、greeter.proto:ct.Ost
		newSvcName := fmt.Sprintf("%s.%s", d.GetPkgName(), svcName)
		data[d.GetFileName()] = append(data[d.GetFileName()], TreeNode{
			NodeID:  newSvcName,
			Type:    ProtoService,
			Data:    svcDef,
			JsonStr: "",
		})
		for _, rpc := range svcDef.Elements {
			rpcData := rpc.(*pp.RPC)
			if rpcData.Parent.(*pp.Service).Name == svcName {
				//ct.Ct:GetName
				data[newSvcName] = append(data[newSvcName], TreeNode{
					NodeID:  rpcData.Name,
					Type:    ProtoMethod,
					Data:    rpcData,
					JsonStr: msgJson[rpcData.RequestType],
				})
				// fd
				fdKey := fmt.Sprintf("%s.%s", newSvcName, rpcData.Name)
				t.ProtoFds[fdKey] = call.ProtoDescriptor{
					FileDescriptor: d.GetFd(),
					Request:        fmt.Sprintf("%s.%s", d.GetPkgName(), rpcData.RequestType),
					Return:         fmt.Sprintf("%s.%s", d.GetPkgName(), rpcData.ReturnsType),
				}
			}
		}
	}
	return data
}

func (t *TreeData) Nodes(uid string) []widget.TreeNodeID {
	var nodes []widget.TreeNodeID
	for _, node := range MenuTree.ProtoData[uid] {
		nodes = append(nodes, node.NodeID)
	}
	return nodes
}

func (t *TreeData) NodeType(uid string) string {
	var _type string
	for _, nodes := range MenuTree.ProtoData {
		for _, nd := range nodes {
			if nd.NodeID == uid {
				_type = nd.Type
				break
			}
		}
	}
	return _type
}

func (t *TreeData) TabName(uid string) string {
	var name, pNodeName string
	for k, nodes := range MenuTree.ProtoData {
		pNodeName = k
		for _, nd := range nodes {
			if nd.NodeID == uid && nd.Type == ProtoMethod {
				name = fmt.Sprintf("%s.%s", pNodeName, nd.NodeID)
				break
			}
		}
	}
	return name
}

func (t *TreeData) RequestJson(uid string) string {
	for _, nodes := range MenuTree.ProtoData {
		for _, nd := range nodes {
			if nd.NodeID == uid {
				return nd.JsonStr
			}
		}
	}
	return ""
}

func (t *TreeData) RemoveAll() {
	t.ProtoData = map[string][]TreeNode{}
	t.ProtoFds = map[string]call.ProtoDescriptor{}
}
