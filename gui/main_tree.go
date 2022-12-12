package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strings"
	"tins-rpc/common"
	theme2 "tins-rpc/theme"
)

func menuTree() *widget.Tree {
	tree := widget.NewTree(
		//childUIDs
		func(uid widget.TreeNodeID) []widget.TreeNodeID {
			return MenuTree.TreeData[uid]
		},
		//isBranch
		func(uid widget.TreeNodeID) bool {
			fmt.Println(uid)
			_, b := MenuTree.TreeData[uid]
			return b
		},
		//create
		func(b bool) fyne.CanvasObject {
			return container.NewHBox(
				widget.NewIcon(theme2.ResourceMSquareIcon),
				widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{}))
		},
		//update
		func(uid widget.TreeNodeID, b bool, object fyne.CanvasObject) {
			if strings.Contains(uid, common.ProtoService) {
				object.(*fyne.Container).Objects[0].(*widget.Icon).SetResource(theme2.ResourceSSquareIcon)
			}
			if strings.Contains(uid, common.ProtoMethod) {
				object.(*fyne.Container).Objects[0].(*widget.Icon).SetResource(theme2.ResourceMSquareIcon)
			}
			object.(*fyne.Container).Objects[1].(*widget.Label).SetText(uid)
		},
	)
	tree.OnSelected = func(uid string) {
		if strings.Contains(uid, "[S]") {
			return
		}
		// 删除 New Tab
		for i, item := range globalWin.tabs.Items {
			if item.Text == "WelCome" {
				globalWin.tabs.RemoveIndex(i)
			}
		}
		// 检测打开并选中
		if _, ok := TabItemList[uid]; ok {
			//设置选中
			globalWin.tabs.Select(TabItemList[uid].TabItem)
			return
		}
		// 添加选项卡
		tabItem := AppendTabItemView(uid, globalWin.tabs)
		// 设置被选中
		tabItem.SelectTree = uid
		// 获取选中方法json
		data := MenuTree.GetRequestJson(uid)
		fmt.Println(uid, data)
		tabItem.RequestText.Text = data
		tabItem.RequestText.Refresh()
		//设置选中
		globalWin.tabs.Select(tabItem.TabItem)
		//保存TabItem
		TabItemList[uid] = tabItem
	}
	tree.OpenAllBranches()
	tree.Refresh()
	return tree
}
