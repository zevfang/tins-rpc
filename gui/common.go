package gui

import (
	"golang.org/x/text/message"
)

func I18n(content string, a ...interface{}) string {
	return message.NewPrinter(Language).Sprintf(content, a...)
}
