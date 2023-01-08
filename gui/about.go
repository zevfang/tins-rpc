package gui

import (
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	tinsTheme "tins-rpc/theme"
)

type about struct {
	aboutDialog dialog.Dialog
}

func newAbout() *about {
	var a about
	content := widget.NewCard("", "", widget.NewRichTextFromMarkdown(I18n(tinsTheme.AboutIntro)))
	a.aboutDialog = dialog.NewCustom(I18n(tinsTheme.AboutTitle), I18n(tinsTheme.ConfirmText), content, globalWin.win)
	return &a
}

func showAbout() {
	newAbout().aboutDialog.Show()
}
