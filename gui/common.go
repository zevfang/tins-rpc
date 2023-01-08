package gui

import (
	"golang.org/x/text/message"
)

func I18n(content string) string {
	return message.NewPrinter(Language).Sprintf(content)
}
