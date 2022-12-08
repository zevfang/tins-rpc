package gui

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"strings"
)

func menuTree() *widget.Tree {
	tree := widget.NewTreeWithStrings(MenuTree.TreeData)
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
	return tree
}
