package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"os"
	"tins-rpc/common"
	theme2 "tins-rpc/theme"
)

type MainWin struct {
	app  fyne.App
	win  fyne.Window
	tabs *container.DocTabs
}

var (
	Version      = "v1.0.0"
	AppID        = "com.tins.rpc.app"
	globalWin    *MainWin
	WindowWidth  float32 = 1400
	WindowHeight float32 = 800
	MenuTree             = *common.NewTreeData()
	TabItemList          = make(map[string]*TabItemView)
)

func init() {
	_ = os.Setenv("FYNE_SCALE", "0.8")
}

func NewMainWin() *MainWin {
	mainWin := new(MainWin)
	// APP
	mainWin.app = app.NewWithID(AppID)
	mainWin.app.Settings().SetTheme(&theme2.LightTheme{})
	mainWin.app.SetIcon(theme2.Ico)

	//WIN
	mainWin.win = mainWin.app.NewWindow(theme2.WindowTitle)
	mainWin.win.Resize(fyne.NewSize(WindowWidth, WindowHeight))
	mainWin.win.SetPadded(false)
	mainWin.win.SetMaster()      //退出窗体则退出程序
	mainWin.win.CenterOnScreen() //屏幕中央

	// PROTO TREE LIST
	tree := menuTree()

	// OPEN FILE
	openBtn := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		fileView := dialog.NewFileOpen(func(file fyne.URIReadCloser, err error) {
			if file == nil {
				return
			}
			filePath := file.URI().Path()
			fileName := file.URI().Name()
			if filePath == "" || fileName == "" {
				return
			}
			err = MenuTree.GetProtoData(filePath)
			if err != nil {
				return
			}
			tree.OpenAllBranches()
			tree.Refresh()
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
	leftHead := container.NewHBox(
		widget.NewLabelWithStyle("Protos", fyne.TextAlignLeading, fyne.TextStyle{}),
		layout.NewSpacer(), openBtn)
	//TODO 线条需要调整一下
	line := canvas.NewRectangle(color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x42})
	line.StrokeWidth = 0.8
	line.SetMinSize(fyne.NewSize(5, 5))

	leftHeadPanel := container.NewVBox(leftHead, line)
	leftPanel := container.NewBorder(leftHeadPanel, nil, nil, nil, tree)

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
