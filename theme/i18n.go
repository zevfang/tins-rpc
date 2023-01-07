package theme

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const Version = "v1.0.2"

var (
	WelComeMsg = "欢迎使用TinsRPC"
	_          = message.SetString(language.English, WelComeMsg, `Welcome to TinsRPC Desktop`)

	WindowTitle = "TinsRPC - RPC客户端请求工具"
	_           = message.SetString(language.English, WindowTitle, `TinsRPC - An rpc client tool`)

	MenuOptTheme = "主题"
	_            = message.SetString(language.English, MenuOptTheme, `Theme`)

	MenuOptThemeDark = "黑暗"
	_                = message.SetString(language.English, MenuOptThemeDark, `Dark`)

	MenuOptThemeLight = "明亮"
	_                 = message.SetString(language.English, MenuOptThemeLight, `Light`)

	ConfirmText = "好"
	_           = message.SetString(language.English, ConfirmText, `OK`)

	SystemTitle = "系统"
	_           = message.SetString(language.English, SystemTitle, `System`)

	OpenFileTitle = "打开文件"
	_             = message.SetString(language.English, OpenFileTitle, `Open File`)

	QuitTitle = "退出"
	_         = message.SetString(language.English, QuitTitle, `Quit`)

	HelpTitle = "帮助"
	_         = message.SetString(language.English, HelpTitle, `Help`)

	AboutTitle = "关于"
	_          = message.SetString(language.English, AboutTitle, `About`)

	AboutIntro = `## TinsRPC 桌面

TinsRPC是一款基于[Fyne](https://fyne.io/)的桌面软件，

这纯粹是出于个人兴趣而建立的。

源代码是 [tins-rpc](https://github.com/zevfang/tins-rpc)。`
	_ = message.SetString(language.English, AboutIntro, `## TinsRPC Desktop  

TinsRPC desktop is a desktop software based on [Fyne](https://fyne.io/),  

which is purely built by personal interests.

The Source code is [tins-rpc](https://github.com/zevfang/tins-rpc).`)

	CheckForUpdatesTitle = "检查更新"
	_                    = message.SetString(language.English, CheckForUpdatesTitle, `Check For Updates…`)

	UpdateYesText = `有可用的新版本。

[tins-rpc %s](%s)`
	_ = message.SetString(language.English, UpdateYesText, `There is a new version available.

[tins-rpc %s](%s)`)

	UpdateNoText = `目前没有可用的更新。`
	_            = message.SetString(language.English, UpdateNoText, `There are currently no updates available.`)

	RunButtonTitle = "运行"
	_              = message.SetString(language.English, RunButtonTitle, `Run`)
)
