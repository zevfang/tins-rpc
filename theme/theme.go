package theme

import (
	_ "embed"
	"fyne.io/fyne/v2"
)

//go:embed PingFang.ttf
var font []byte

var myfont = &fyne.StaticResource{
	StaticName:    "PingFang.ttf",
	StaticContent: font,
}

//func (m *MyTheme) Font(_ fyne.TextStyle) fyne.Resource {
//	return myfont
//	//return theme.DefaultTheme().Font(s)
//}

//go:embed icon.png
var logo []byte

var Ico = &fyne.StaticResource{
	StaticName:    "icon.png",
	StaticContent: logo,
}
