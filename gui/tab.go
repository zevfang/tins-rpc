package gui

import (
	"encoding/json"
	"fmt"
	"image/color"
	"strings"
	"time"

	"tins-rpc/call"
	"tins-rpc/common"
	"tins-rpc/store"
	tinsTheme "tins-rpc/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const centerPanelOffset = 0.96

type TabItemView struct {
	UrlInput           *widget.Entry
	RequestText        *widget.Entry
	ResponseText       *widget.Entry
	MetadataText       *widget.Entry
	MetadataTextHeight float32
	FrameSelect        *widget.Select
	UsedTimeLabel      *widget.Label
	CallButton         *widget.Button
	CodeText           *widget.Entry //浏览proto源代码
	SelectTree         string
	ProtoName          string
	TabItem            *container.TabItem
	contentPanel       *fyne.Container
	metadataPanel      *fyne.Container
	centerPanel        *container.Split
}

func AppendTabItemView(tabTitle string, tabs *container.DocTabs) *TabItemView {
	tabItemView := &TabItemView{}
	tabItemView.ProtoName = fmt.Sprintf("%s.proto", strings.Split(tabTitle, ".")[0]) //本地存储使用

	// URI TEXT
	tabItemView.UrlInput = widget.NewEntry()
	tabItemView.UrlInput.PlaceHolder = "127.0.0.1:8080"

	// 本地获取url
	storeUrl := StoreData.Url.Get(tabItemView.ProtoName)
	if len(storeUrl.Url) > 0 {
		tabItemView.UrlInput.Text = storeUrl.Url
	}
	tabItemView.UrlInput.Validator = validation.NewRegexp(`\S+`, "URL must not be empty")
	tabItemView.UrlInput.Validator = validation.NewRegexp(`((\d{1,3}.){3}\d{1,3}:\d+)`, "please input right URL")

	// REQ TEXT
	tabItemView.RequestText = widget.NewMultiLineEntry()
	tabItemView.RequestText.PlaceHolder = "Editor(json)"

	//tabItemView.RequestText.TappedSecondary()
	// RESP TEXT
	tabItemView.ResponseText = widget.NewMultiLineEntry()
	tabItemView.ResponseText.PlaceHolder = "Response(json)"

	// METADATA TEXT
	tabItemView.newMetadataContainer()

	// CALL BUTTON
	tabItemView.CallButton = widget.NewButtonWithIcon(I18n(tinsTheme.RunButtonTitle), tinsTheme.ResourceRunIcon, func() {
		tabItemView.OnCall()
	})

	// 框架选项
	tabItemView.FrameSelect = widget.NewSelect(call.FrameTypes, func(s string) {})
	if len(storeUrl.Frame) == 0 {
		tabItemView.FrameSelect.SetSelected(call.RPCX.ToString())
	} else {
		tabItemView.FrameSelect.SetSelected(storeUrl.Frame)
	}

	// 耗时显示
	tabItemView.UsedTimeLabel = widget.NewLabel("")

	headPanel := container.NewGridWithColumns(6,
		tabItemView.UrlInput,
		tabItemView.FrameSelect,
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(),
			container.NewGridWithColumns(2, tabItemView.UsedTimeLabel, tabItemView.CallButton),
		))
	tabItemView.centerPanel = container.NewVSplit(
		container.NewHSplit(tabItemView.RequestText, tabItemView.ResponseText),
		tabItemView.metadataPanel)
	tabItemView.centerPanel.SetOffset(centerPanelOffset)

	tabItemView.contentPanel = container.NewBorder(headPanel, nil, nil, nil, tabItemView.centerPanel)
	tabItemView.TabItem = container.NewTabItem(tabTitle, tabItemView.contentPanel)
	tabs.Append(tabItemView.TabItem)
	return tabItemView
}

func (tiv *TabItemView) newMetadataContainer() {
	tiv.MetadataText = widget.NewMultiLineEntry()
	tiv.MetadataText.Hide()
	metadataTxt := canvas.NewText("  ▲ METADATA", color.NRGBA{R: 0x21, G: 0x96, B: 0xf3, A: 0xff})
	metadataTxt.Alignment = fyne.TextAlignLeading
	metadataTxt.TextSize = 14
	metadataBtn := widget.NewButton("", func() {
		if tiv.MetadataText.Visible() {
			metadataTxt.Text = "  ▲ METADATA"
			tiv.MetadataText.Hide()
		} else {
			metadataTxt.Text = "  ▼ METADATA"
			tiv.MetadataText.Show()
		}
		tiv.centerPanel.SetOffset(centerPanelOffset)
		tiv.centerPanel.Refresh()
	})
	metadataPanelTop := container.NewMax(metadataBtn, metadataTxt)
	metadataPanelBottom := container.NewMax(tiv.MetadataText)
	tiv.metadataPanel = container.NewVBox(
		metadataPanelTop,
		metadataPanelBottom,
	)
	// 监听panel重置txt宽度
	go func() {
		time.Sleep(1 * time.Second)
		for {
			_centerPanelSize := tiv.centerPanel.Trailing.Size()
			centerPanelHeight := _centerPanelSize.Height

			_metadataPanelTopSize := metadataPanelTop.Size()
			metadataPanelTopHeight := _metadataPanelTopSize.Height

			metadataTextHeight := centerPanelHeight - metadataPanelTopHeight
			tiv.MetadataText.Resize(fyne.NewSize(tiv.MetadataText.Size().Width, metadataTextHeight))
			tiv.MetadataText.Refresh()
			time.Sleep(200 * time.Millisecond)
		}
	}()
}

func AppendTabItemCodeView(tabTitle string, protoBody string, tabs *container.DocTabs) *TabItemView {
	tabItemCodeView := &TabItemView{}
	tabItemCodeView.CodeText = widget.NewMultiLineEntry()
	tabItemCodeView.CodeText.Text = protoBody
	contentPanel := container.NewBorder(nil, nil, nil, nil, tabItemCodeView.CodeText)
	tabItemCodeView.TabItem = container.NewTabItem(tabTitle, contentPanel)
	tabs.Append(tabItemCodeView.TabItem)
	return tabItemCodeView
}

func (t *TabItemView) OnCall() {
	if len(t.UrlInput.Text) == 0 {
		t.ResponseText.Text = "URL not found"
		t.ResponseText.Refresh()
		return
	}
	if len(t.RequestText.Text) == 0 {
		t.RequestText.Text = "Request not found"
		t.RequestText.Refresh()
		return
	}
	address := t.UrlInput.Text
	url := strings.Split(address, ":")
	if len(url) != 2 {
		t.ResponseText.Text = "URL failed,Does not include[:]."
		t.ResponseText.Refresh()
		return
	}
	if strings.Contains(address, "/") {
		t.ResponseText.Text = "URL failed,no path required."
		t.ResponseText.Refresh()
		return
	}

	fmt.Println("框架：", t.FrameSelect.Selected)
	svcPath := strings.Split(t.SelectTree, ".")
	fmt.Println("服务：", t.SelectTree)
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
	// 保存url信息到本地
	protoName := fmt.Sprintf("%s.proto", svcPath[0])
	StoreData.Url.Set(protoName, store.UrlStoreModel{
		Url:   t.UrlInput.Text,
		Frame: t.FrameSelect.Selected,
	})

	//禁用按钮
	t.CallButton.Disable()

	go func() {
		tms := time.Now()
		_, body, err := call.Call(t.FrameSelect.Selected, call.RequestData{
			Fd:            MenuTree.ProtoFds[t.SelectTree],
			Address:       address,
			PackageName:   svcPath[0],
			ServicePath:   svcPath[1],
			ServiceMethod: svcPath[2],
			Metadata:      metadata,
			Payload:       payload,
		})
		if err != nil {
			t.ResponseText.Text = err.Error()
			t.ResponseText.Refresh()
			// 显示耗时
			t.UsedTimeLabel.SetText(fmt.Sprintf("%d ms", time.Since(tms).Milliseconds()))
			t.UsedTimeLabel.Refresh()
			// 启用按钮
			t.CallButton.Enable()
			return
		}
		t.ResponseText.Text = common.FormatJSON(body)
		t.ResponseText.Refresh()
		// 显示耗时
		t.UsedTimeLabel.SetText(fmt.Sprintf("%d ms", time.Since(tms).Milliseconds()))
		t.UsedTimeLabel.Refresh()
		// 启用按钮
		t.CallButton.Enable()
	}()

}
