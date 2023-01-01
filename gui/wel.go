package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	theme2 "tins-rpc/theme"
)

func initWelcome() *container.TabItem {
	var welcomeTabItem *container.TabItem
	logo := canvas.NewImageFromResource(theme2.ResourceLogoIcon)
	logo.FillMode = canvas.ImageFillContain
	logo.SetMinSize(fyne.NewSize(362*0.8, 192*0.8))

	welcomeTabItem = container.NewTabItemWithIcon("WelCome", nil, nil)
	wel := widget.NewRichTextFromMarkdown("# " + theme2.WelComeMsg)
	for i := range wel.Segments {
		if seg, ok := wel.Segments[i].(*widget.TextSegment); ok {
			seg.Style.Alignment = fyne.TextAlignCenter
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
