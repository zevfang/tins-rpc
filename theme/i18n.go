package theme

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const Version = "v1.0.2"

var (
	WelComeMsg = "Welcome to TinsRPC Client"
	_          = message.SetString(language.Chinese, WelComeMsg, `欢迎使用TinsRPC(罐头)客户端`)

	WelComeTabTitle = "Welcome"
	_               = message.SetString(language.Chinese, WelComeTabTitle, `欢迎`)

	WindowTitle = "TinsRPC"
	_           = message.SetString(language.Chinese, WindowTitle, `TinsRPC(罐头)`)

	MenuOptTheme = "Theme"
	_            = message.SetString(language.Chinese, MenuOptTheme, `主题`)

	MenuOptThemeDark = "Dark"
	_                = message.SetString(language.Chinese, MenuOptThemeDark, `黑暗`)

	MenuOptThemeLight = "Light"
	_                 = message.SetString(language.Chinese, MenuOptThemeLight, `明亮`)

	MenuOptLanguage = "Language"
	_               = message.SetString(language.Chinese, MenuOptLanguage, `语言`)

	MenuOptLanguageZhCN = "Chinese"
	_                   = message.SetString(language.Chinese, MenuOptLanguageZhCN, `中文`)

	MenuOptLanguageEnUS = "English"
	_                   = message.SetString(language.Chinese, MenuOptLanguageEnUS, "英文")

	ConfirmText = "OK"
	_           = message.SetString(language.Chinese, ConfirmText, `好`)

	SystemTitle = "System"
	_           = message.SetString(language.Chinese, SystemTitle, `系统`)

	OpenFileTitle = "Open File"
	_             = message.SetString(language.Chinese, OpenFileTitle, `打开文件`)

	QuitTitle = "Quit"
	_         = message.SetString(language.Chinese, QuitTitle, `退出`)

	HelpTitle = "Help"
	_         = message.SetString(language.Chinese, HelpTitle, `帮助`)

	AboutTitle = "About"
	_          = message.SetString(language.Chinese, AboutTitle, `关于`)

	AboutIntro = `## TinsRPC %s RPC Client  

TinsRPC desktop is a desktop software based on [Fyne](https://fyne.io/),

The source code is [tins-rpc](https://github.com/zevfang/tins-rpc).`
	_ = message.SetString(language.Chinese, AboutIntro, `## TinsRPC %s 客户端工具

TinsRPC是一款基于[Fyne](https://fyne.io/)的桌面软件，

源代码仓库 [tins-rpc](https://github.com/zevfang/tins-rpc)。`)

	CheckForUpdatesTitle = "Check For Updates…"
	_                    = message.SetString(language.Chinese, CheckForUpdatesTitle, `检查更新`)

	UpdateYesText = `There is a new version available.

[tins-rpc %s](%s)`
	_ = message.SetString(language.Chinese, UpdateYesText, `有可用的新版本。

[tins-rpc %s](%s)`)

	UpdateNoText = `There are currently no updates available.`
	_            = message.SetString(language.Chinese, UpdateNoText, `目前没有可用的更新`)

	RunButtonTitle = "Run"
	_              = message.SetString(language.Chinese, RunButtonTitle, `运行`)

	PromptTitle = "Prompt"
	_           = message.SetString(language.Chinese, PromptTitle, `提示`)

	PromptRestartContentText = "Take effect after restarting the application"
	_                        = message.SetString(language.Chinese, PromptRestartContentText, `重启应用后生效`)
)
