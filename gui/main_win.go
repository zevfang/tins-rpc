package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"os"
	"tins-rpc/data"
	theme2 "tins-rpc/theme"
)

type MainWin struct {
	app  fyne.App
	win  fyne.Window
	tabs *container.DocTabs
	tree *widget.Tree
}

var (
	Version      = "v1.0.0"
	AppID        = "com.tins.call.app"
	globalWin    *MainWin
	WindowWidth  float32 = 1400
	WindowHeight float32 = 800
	MenuTree             = NewTreeData()
	TabItemList          = make(map[string]*TabItemView)
	StorageData          = data.NewStorage()
)

func init() {
	_ = os.Setenv("FYNE_SCALE", "0.8")
}

func NewMainWin() *MainWin {
	mainWin := new(MainWin)
	// APP
	mainWin.app = app.NewWithID(AppID)
	mainWin.app.Settings().SetTheme(&theme2.LightTheme{})
	mainWin.app.SetIcon(theme2.ResourceLogoIcon)

	//WIN
	mainWin.win = mainWin.app.NewWindow(theme2.WindowTitle)
	mainWin.win.Resize(fyne.NewSize(WindowWidth, WindowHeight))
	mainWin.win.SetPadded(false)
	mainWin.win.SetMaster()      //退出窗体则退出程序
	mainWin.win.CenterOnScreen() //屏幕中央

	//MenuTree.TreeData = map[string][]string{
	//	"":         {"ct.proto"},
	//	"ct.proto": {"ct.Ct"},
	//	"ct.Ct":    {"Ct.GetTable", "Ct.GetTopic"},
	//}
	// PROTO TREE LIST
	mainWin.tree = menuTree()

	// Refresh
	refreshBtn := widget.NewButtonWithIcon("", theme2.ResourceRefreshIcon, func() {

	})
	// CLEAR
	clearBtn := widget.NewButtonWithIcon("", theme2.ResourceClearIcon, func() {
		MenuTree.RemoveAll()
		mainWin.tree.Refresh()
	})
	// OPEN FILE
	openBtn := widget.NewButtonWithIcon("", theme2.ResourceAddIcon, func() {
		fileView := dialog.NewFileOpen(func(file fyne.URIReadCloser, err error) {
			if file == nil {
				return
			}
			filePath := file.URI().Path()
			fileName := file.URI().Name()
			if filePath == "" || fileName == "" {
				return
			}
			err = MenuTree.Append(filePath)
			if err != nil {
				return
			}
			mainWin.tree.OpenAllBranches()
			mainWin.tree.Refresh()
		}, mainWin.win)
		fileView.SetFilter(storage.NewExtensionFileFilter([]string{".proto"}))
		fileView.Resize(fyne.Size{
			Width:  700,
			Height: 550,
		})
		fileView.Show()
	})

	//TABS
	welcomeTab := initWelcome()
	mainWin.tabs = container.NewDocTabs(welcomeTab)
	mainWin.tabs.OnClosed = func(item *container.TabItem) {
		if len(mainWin.tabs.Items) == 0 {
			mainWin.tabs.Append(welcomeTab)
			TabItemList = make(map[string]*TabItemView) //清空tabItem记录
		}
		delete(TabItemList, item.Text)
	}

	// LEFT
	//line := canvas.NewLine(color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x42})
	//line.StrokeWidth = 0.8
	leftBtnBox := container.NewHBox(refreshBtn, clearBtn, layout.NewSpacer(), openBtn)
	//container.NewVBox(line, leftBtnBox)
	leftHeadCard := widget.NewCard("", "", leftBtnBox)
	leftPanel := container.NewBorder(leftHeadCard, nil, nil, nil, mainWin.tree)

	// CONTENT
	content := container.NewHSplit(leftPanel, mainWin.tabs)
	content.SetOffset(0.25)

	home := container.NewBorder(nil, nil, nil, nil, content)
	mainWin.win.SetContent(home)
	mainWin.win.SetMainMenu(mainMenu())

	globalWin = mainWin
	return globalWin
}

func (m *MainWin) Run() {
	m.win.ShowAndRun()
}
