package theme

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const Version = "v1.0.2"

var (
	WelComeMsg = "欢迎使用TinsRPC"
	_          = message.SetString(language.English, WelComeMsg, `Welcome to TinsRPC Desktop`)

	WelComeTabTitle = "欢迎"
	_               = message.SetString(language.English, WelComeTabTitle, `Welcome`)

	WindowTitle = "TinsRPC - RPC客户端工具"
	_           = message.SetString(language.English, WindowTitle, `TinsRPC - An rpc client tool`)

	MenuOptTheme = "主题"
	_            = message.SetString(language.English, MenuOptTheme, `Theme`)

	MenuOptThemeDark = "黑暗"
	_                = message.SetString(language.English, MenuOptThemeDark, `Dark`)

	MenuOptThemeLight = "明亮"
	_                 = message.SetString(language.English, MenuOptThemeLight, `Light`)

	MenuOptLanguage = "语言"
	_               = message.SetString(language.English, MenuOptLanguage, `Language`)

	MenuOptLanguageZhCN = "中文"
	_                   = message.SetString(language.English, MenuOptLanguageZhCN, `Chinese`)

	MenuOptLanguageEnUS = "英文"
	_                   = message.SetString(language.English, MenuOptLanguageEnUS, `English`)

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

	AboutIntro = `## TinsRPC 客户端工具

TinsRPC是一款基于[Fyne](https://fyne.io/)的桌面软件，

源代码仓库 [tins-rpc](https://github.com/zevfang/tins-rpc)。`
	_ = message.SetString(language.English, AboutIntro, `## TinsRPC client tool  

TinsRPC desktop is a desktop software based on [Fyne](https://fyne.io/),

The source code is [tins-rpc](https://github.com/zevfang/tins-rpc).`)

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

	PromptTitle = "提示"
	_           = message.SetString(language.English, PromptTitle, `Prompt`)

	PromptRestartContentText = "重启应用后生效"
	_                        = message.SetString(language.English, PromptRestartContentText, `Take effect after restarting the application`)
)
