package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"tins-rpc/common"
	tinsTheme "tins-rpc/theme"
)

func mainMenu() *fyne.MainMenu {
	_themeMenu := themeMenu()

	sysMenu := &fyne.Menu{
		Label: tinsTheme.SystemTitle,
		Items: []*fyne.MenuItem{
			{Label: tinsTheme.OpenFileTitle, Action: OpenFileAction},
			_themeMenu,
			fyne.NewMenuItemSeparator(),
			{Label: tinsTheme.QuitTitle, IsQuit: true, Action: QuitAction},
		},
	}

	helpMenu := &fyne.Menu{
		Label: tinsTheme.HelpTitle,
		Items: []*fyne.MenuItem{
			{Label: tinsTheme.AboutTitle, Action: AboutAction},
			{Label: tinsTheme.CheckForUpdatesTitle, Action: CheckForUpdateAction},
		},
	}
	return fyne.NewMainMenu(sysMenu, helpMenu)
}

func OpenFileAction() {

}

func QuitAction() {

}

func AboutAction() {
	showAbout()
}

func themeMenu() *fyne.MenuItem {
	var (
		themeOpt   *fyne.MenuItem
		themeDark  *fyne.MenuItem
		themeLight *fyne.MenuItem
	)
	// Option-Theme
	themeDark = fyne.NewMenuItem(tinsTheme.MenuOptThemeDark, func() {
		globalWin.app.Settings().SetTheme(tinsTheme.DarkTheme{})
		themeDark.Checked = true
		themeLight.Checked = false
		_ = globalConfig.Theme.Set("__DARK__")
		globalWin.mainMenu.Refresh()
	})
	themeDark.Checked = true
	themeLight = fyne.NewMenuItem(tinsTheme.MenuOptThemeLight, func() {
		globalWin.app.Settings().SetTheme(tinsTheme.LightTheme{})
		themeDark.Checked = false
		themeLight.Checked = true
		_ = globalConfig.Theme.Set("__LIGHT__")
		globalWin.mainMenu.Refresh()
	})
	themeLight.Checked = true
	t, _ := globalConfig.Theme.Get()
	if t == "__DARK__" {
		themeLight.Checked = false
	} else {
		themeDark.Checked = false
	}

	themeOpt = fyne.NewMenuItem(tinsTheme.MenuOptTheme, nil)
	themeOpt.ChildMenu = fyne.NewMenu("",
		themeDark,
		themeLight,
	)
	return themeOpt
}

func CheckForUpdateAction() {
	var content *widget.Card
	isUpdate, tagName, tagUrl := common.CheckForUpdates()
	if isUpdate {
		content = widget.NewCard("", "", widget.NewRichTextFromMarkdown(fmt.Sprintf(tinsTheme.UpdateYesText, tagName, tagUrl)))
	} else {
		content = widget.NewCard("", "", widget.NewRichTextFromMarkdown(tinsTheme.UpdateNoText))
	}
	updateDialog := dialog.NewCustom(tinsTheme.CheckForUpdatesTitle, tinsTheme.ConfirmText, content, globalWin.win)
	updateDialog.Show()
}
