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
	sysMenu := &fyne.Menu{
		Label: "System",
		Items: []*fyne.MenuItem{
			{Label: "Open File", Action: OpenFileAction},
			fyne.NewMenuItemSeparator(),
			{Label: "Quit", IsQuit: true, Action: QuitAction},
		},
	}

	helpMenu := &fyne.Menu{
		Label: "Help",
		Items: []*fyne.MenuItem{
			{Label: "About", Action: AboutAction},
			{Label: "Check For Updates…", Action: CheckForUpdateAction},
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
