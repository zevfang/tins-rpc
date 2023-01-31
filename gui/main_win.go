package gui

import (
	"fmt"
	"os"

	"tins-rpc/store"
	tinsTheme "tins-rpc/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/text/language"
)

type MainWin struct {
	app      fyne.App
	win      fyne.Window
	tabs     *container.DocTabs
	tree     *widget.Tree
	mainMenu *fyne.MainMenu
}

var (
	AppID        = "com.tins.rpc.app"
	globalWin    *MainWin
	globalConfig *tinsTheme.Config
	WindowWidth  float32 = 1400
	WindowHeight float32 = 800
	MenuTree             = NewTreeData()
	TabItemList          = make(map[string]*TabItemView)
	StoreData            = store.InitStore()
	Language     language.Tag
)

func init() {
	_ = os.Setenv("FYNE_SCALE", "0.8")
}

func NewMainWin() *MainWin {
	mainWin := new(MainWin)
	globalWin = mainWin
	// APP
	mainWin.app = app.NewWithID(AppID)
	//mainWin.app.Settings().SetTheme(&tinsTheme.DarkTheme{})
	mainWin.app.SetIcon(tinsTheme.ResourceLogoIcon)
	globalConfig = tinsTheme.NewConfig()
	// Theme
	preTheme, _ := globalConfig.Theme.Get()
	switch preTheme {
	case "__DARK__":
		mainWin.app.Settings().SetTheme(&tinsTheme.DarkTheme{})
	case "__LIGHT__":
		mainWin.app.Settings().SetTheme(&tinsTheme.LightTheme{})
	}
	// Language
	preLanguage, _ := globalConfig.Language.Get()
	switch preLanguage {
	case "__en-US__":
		Language = language.English
	case "__zh-CN__":
		Language = language.Chinese
	}

	//WIN
	mainWin.win = mainWin.app.NewWindow(I18n(tinsTheme.WindowTitle))
	mainWin.win.Resize(fyne.NewSize(WindowWidth, WindowHeight))
	mainWin.win.SetPadded(false)
	mainWin.win.SetMaster()      //退出窗体则退出程序
	mainWin.win.CenterOnScreen() //屏幕中央

	// PROTO TREE LIST
	mainWin.tree = menuTree()

	// Refresh
	refreshBtn := widget.NewButtonWithIcon("", tinsTheme.ResourceRefreshIcon, func() {
		MenuTree.RefreshAll()
		mainWin.tree.Refresh()
	})
	// CLEAR
	clearBtn := widget.NewButtonWithIcon("", tinsTheme.ResourceClearIcon, func() {
		MenuTree.RemoveAll()
		mainWin.tree.Refresh()
	})
	// OPEN FILE
	openBtn := widget.NewButtonWithIcon("", tinsTheme.ResourceAddIcon, func() {
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
	leftBtnBox := container.NewHBox(clearBtn, refreshBtn, layout.NewSpacer(), openBtn)
	leftCard := container.NewBorder(leftBtnBox, nil, nil, nil, mainWin.tree) //边框
	leftPanel := widget.NewCard("", "", leftCard)

	// CONTENT
	content := container.NewHSplit(leftPanel, mainWin.tabs)
	content.SetOffset(0.25)

	home := container.NewBorder(nil, nil, nil, nil, content)
	mainWin.win.SetContent(home)
	mainWin.mainMenu = mainMenu()
	mainWin.win.SetMainMenu(mainWin.mainMenu)

	return globalWin
}

func (m *MainWin) Run() {
	m.win.ShowAndRun()
	m.win.SetCloseIntercept(func() {
		//TODO 这里是否可以重启程序？
		fmt.Println("close")
	})
}
