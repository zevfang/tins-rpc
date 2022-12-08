package gui

import "fyne.io/fyne/v2"

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
