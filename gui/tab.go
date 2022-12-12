package gui

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strings"
	"tins-rpc/common"
	theme2 "tins-rpc/theme"
)

type TabItemView struct {
	UriInput     *widget.Entry
	RequestText  *widget.Entry
	ResponseText *widget.Entry
	MetadataText *widget.Entry
	RpcSelect    *widget.Select
	CallButton   *widget.Button
	SelectTree   string
	TabItem      *container.TabItem
}

func AppendTabItemView(tabTitle string, tabs *container.DocTabs) *TabItemView {
	tabItemView := &TabItemView{}
	// URI TEXT
	tabItemView.UriInput = widget.NewEntry()
	tabItemView.UriInput.PlaceHolder = "127.0.0.1:8080"

	// REQ TEXT
	tabItemView.RequestText = widget.NewMultiLineEntry()
	tabItemView.RequestText.PlaceHolder = "Editor(json)"

	//tabItemView.RequestText.TappedSecondary()
	// RESP TEXT
	tabItemView.ResponseText = widget.NewMultiLineEntry()
	tabItemView.ResponseText.PlaceHolder = "Response(json)"

	// METADATA TEXT
	tabItemView.MetadataText = widget.NewMultiLineEntry()
	tabItemView.MetadataText.SetPlaceHolder("METADATA")

	// CALL BUTTON
	tabItemView.CallButton = widget.NewButtonWithIcon("Run", theme2.ResourceRunIcon, func() {
		tabItemView.OnCall()
	})

	// 框架选项
	tabItemView.RpcSelect = widget.NewSelect([]string{common.RPCX}, func(s string) {})
	tabItemView.RpcSelect.SetSelected(common.RPCX)

	headPanel := container.NewGridWithColumns(6,
		tabItemView.UriInput,
		tabItemView.RpcSelect,
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		container.NewHBox(layout.NewSpacer(), tabItemView.CallButton))

	centerPanel := container.NewVSplit(
		container.NewHSplit(tabItemView.RequestText, tabItemView.ResponseText),
		tabItemView.MetadataText)
	centerPanel.SetOffset(0.9)

	contentPanel := container.NewBorder(headPanel, nil, nil, nil, centerPanel)
	tabItemView.TabItem = container.NewTabItem(tabTitle, contentPanel)
	tabs.Append(tabItemView.TabItem)
	return tabItemView
}

func (t *TabItemView) OnCall() {
	if len(t.UriInput.Text) == 0 {
		t.ResponseText.Text = "URI not found"
		t.ResponseText.Refresh()
		return
	}
	if len(t.RequestText.Text) == 0 {
		t.RequestText.Text = "Request not found"
		t.RequestText.Refresh()
		return
	}
	address := t.UriInput.Text
	uri := strings.Split(address, ":")
	if len(uri) != 2 {
		t.ResponseText.Text = "URI failed,Does not include[:]."
		t.ResponseText.Refresh()
		return
	}
	if strings.Contains(address, "/") {
		t.ResponseText.Text = "URI failed,no path required."
		t.ResponseText.Refresh()
		return
	}
	if strings.Contains(t.SelectTree, "[S]") {
		t.ResponseText.Text = "Please select a method."
		t.ResponseText.Refresh()
		return
	}
	fmt.Println("框架：", t.RpcSelect.Selected)
	serviceAndRpc := strings.Split(strings.Split(t.SelectTree, " ")[1], ".")
	fmt.Println("服务：", serviceAndRpc)
	payload := []byte(t.RequestText.Text)

	metadata := make(map[string]string)
	metadataText := t.MetadataText.Text
	if len(metadataText) > 0 {
		err := json.Unmarshal([]byte(metadataText), &metadata)
		if err != nil {
			t.ResponseText.Text = fmt.Sprintf("metadata error,%v", err.Error())
			t.ResponseText.Refresh()
			return
		}
	}

	go func() {
		_, body, err := common.Call(t.RpcSelect.Selected, common.RequestData{
			Address:       address,
			ServicePath:   serviceAndRpc[0],
			ServiceMethod: serviceAndRpc[1],
			Metadata:      metadata,
			Payload:       payload,
		})
		if err != nil {
			t.ResponseText.Text = err.Error()
			t.ResponseText.Refresh()
			return
		}
		t.ResponseText.Text = common.ParseJSON(body)
		t.ResponseText.Refresh()
	}()

}
