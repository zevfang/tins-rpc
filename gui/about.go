package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"tins-rpc/common"
	tinsTheme "tins-rpc/theme"
)

type about struct {
	aboutDialog dialog.Dialog
}

func newAbout() *about {
	var a about
	content := widget.NewCard("", "", widget.NewRichTextFromMarkdown(I18n(tinsTheme.AboutIntro, common.Version)))
	a.aboutDialog = dialog.NewCustom(I18n(tinsTheme.AboutTitle), I18n(tinsTheme.ConfirmText), content, globalWin.win)
	return &a
}

func showAbout() {
	newAbout().aboutDialog.Show()
	//runPopUp(globalWin.win)
}

func runPopUp(w fyne.Window) {
	var modal *widget.PopUp
	modal = widget.NewModalPopUp(
		container.NewVBox(
			widget.NewLabel("bar"),
			widget.NewButton("Close", func() { modal.Hide() }),
		),
		w.Canvas(),
	)
	modal.Resize(fyne.NewSize(WindowWidth/2, WindowHeight))
	modal.Move(fyne.Position{250, 250})
	modal.Show()
}
