package gui

import (
	tinsTheme "tins-rpc/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func initWelcome() *container.TabItem {
	var welcomeTabItem *container.TabItem
	logo := canvas.NewImageFromResource(tinsTheme.ResourceLogoIcon)
	logo.FillMode = canvas.ImageFillContain
	logo.SetMinSize(fyne.NewSize(362*0.8, 192*0.8))

	welcomeTabItem = container.NewTabItemWithIcon(I18n(tinsTheme.WelComeTabTitle), nil, nil)
	wel := widget.NewRichTextFromMarkdown("# " + I18n(tinsTheme.WelComeMsg))
	for i := range wel.Segments {
		if seg, ok := wel.Segments[i].(*widget.TextSegment); ok {
			seg.Style.Alignment = fyne.TextAlignCenter
			seg.Style.ColorName = theme.ColorNamePrimary
		}
	}

	shortCuts := widget.NewForm()
	appRegister := make(map[string]string)
	//appRegister["快捷键1"] = "Ctrl+C"
	//appRegister["快捷键2"] = "Ctrl+V"
	for k, myApp := range appRegister {
		shortCuts.Append(k, widget.NewLabelWithStyle(myApp, fyne.TextAlignCenter, fyne.TextStyle{}))
	}
	//shortCuts.Append("Show/Hide", new_widget.NewLabelWithStyle("This is a call client tool!", fyne.TextAlignCenter, fyne.TextStyle{}))

	welcomeTabItem.Content = container.NewCenter(
		container.NewVBox(
			wel,
			logo,
			shortCuts,
		))

	return welcomeTabItem
}
