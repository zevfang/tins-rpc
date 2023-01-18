package gui

import (
	"tins-rpc/common"
	tinsTheme "tins-rpc/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/text/language"
)

func mainMenu() *fyne.MainMenu {
	_themeMenu := themeMenu()
	_i18nMenu := i18nMenu()
	sysMenu := &fyne.Menu{
		Label: I18n(tinsTheme.SystemTitle),
		Items: []*fyne.MenuItem{
			{Label: I18n(tinsTheme.OpenFileTitle), Action: OpenFileAction},
			_themeMenu,
			_i18nMenu,
			fyne.NewMenuItemSeparator(),
			{Label: I18n(tinsTheme.QuitTitle), IsQuit: true, Action: QuitAction},
		},
	}

	helpMenu := &fyne.Menu{
		Label: I18n(tinsTheme.HelpTitle),
		Items: []*fyne.MenuItem{
			{Label: I18n(tinsTheme.AboutTitle), Action: AboutAction},
			{Label: I18n(tinsTheme.CheckForUpdatesTitle), Action: CheckForUpdateAction},
		},
	}
	return fyne.NewMainMenu(sysMenu, helpMenu)
}

func themeMenu() *fyne.MenuItem {
	var (
		themeOpt   *fyne.MenuItem
		themeDark  *fyne.MenuItem
		themeLight *fyne.MenuItem
	)
	// Option-Theme
	themeDark = fyne.NewMenuItem(I18n(tinsTheme.MenuOptThemeDark), func() {
		globalWin.app.Settings().SetTheme(tinsTheme.DarkTheme{})
		themeDark.Checked = true
		themeLight.Checked = false
		_ = globalConfig.Theme.Set("__DARK__")
		globalWin.mainMenu.Refresh()
	})
	themeDark.Checked = true
	themeLight = fyne.NewMenuItem(I18n(tinsTheme.MenuOptThemeLight), func() {
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

	themeOpt = fyne.NewMenuItem(I18n(tinsTheme.MenuOptTheme), nil)
	themeOpt.ChildMenu = fyne.NewMenu("",
		themeDark,
		themeLight,
	)
	return themeOpt
}

func i18nMenu() *fyne.MenuItem {
	dialogRestartTip := func() dialog.Dialog {
		content := widget.NewCard("", "", widget.NewRichTextFromMarkdown(I18n(tinsTheme.PromptRestartContentText)))
		return dialog.NewCustom(I18n(tinsTheme.PromptTitle), I18n(tinsTheme.ConfirmText), content, globalWin.win)
	}
	var (
		i18nOpt  *fyne.MenuItem
		i18nEnUs *fyne.MenuItem
		i18nZhCn *fyne.MenuItem
	)
	// en-US
	i18nEnUs = fyne.NewMenuItem(I18n(tinsTheme.MenuOptLanguageEnUS), func() {
		i18nEnUs.Checked = true
		i18nZhCn.Checked = false
		_ = globalConfig.Language.Set("__en-US__")
		globalWin.mainMenu.Refresh()
		dialogRestartTip().Show()
		Language = language.English
	})
	i18nEnUs.Checked = true
	// zh-CN
	i18nZhCn = fyne.NewMenuItem(I18n(tinsTheme.MenuOptLanguageZhCN), func() {
		i18nEnUs.Checked = false
		i18nZhCn.Checked = true
		_ = globalConfig.Language.Set("__zh-CN__")
		globalWin.mainMenu.Refresh()
		dialogRestartTip().Show()
		Language = language.Chinese
	})
	i18nZhCn.Checked = true
	t, _ := globalConfig.Language.Get()
	if t == "__zh-CN__" {
		i18nEnUs.Checked = false
	} else {
		i18nZhCn.Checked = false
	}

	i18nOpt = fyne.NewMenuItem(I18n(tinsTheme.MenuOptLanguage), nil)
	i18nOpt.ChildMenu = fyne.NewMenu("",
		i18nEnUs,
		i18nZhCn,
	)
	return i18nOpt
}

func OpenFileAction() {

}

func QuitAction() {
	globalWin.app.Quit()
}

func AboutAction() {
	showAbout()
}

func CheckForUpdateAction() {
	var content *widget.Card
	isUpdate, tagName, tagUrl := common.CheckForUpdates()
	if isUpdate {
		content = widget.NewCard("", "", widget.NewRichTextFromMarkdown(I18n(tinsTheme.UpdateYesText, tagName, tagUrl)))
	} else {
		content = widget.NewCard("", "", widget.NewRichTextFromMarkdown(I18n(tinsTheme.UpdateNoText)))
	}
	updateDialog := dialog.NewCustom(I18n(tinsTheme.CheckForUpdatesTitle), I18n(tinsTheme.ConfirmText), content, globalWin.win)
	updateDialog.Show()
}
