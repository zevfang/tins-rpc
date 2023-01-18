package theme

import (
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

type Config struct {
	Theme    binding.String
	Language binding.String
}

func NewConfig() *Config {
	c := &Config{
		Theme:    binding.NewString(),
		Language: binding.NewString(),
	}

	_ = c.Theme.Set(fyne.CurrentApp().Preferences().StringWithFallback("theme", "__DARK__"))
	_ = c.Language.Set(fyne.CurrentApp().Preferences().StringWithFallback("language", "__en-US__"))

	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(c)
	s := reflect.ValueOf(c).Elem()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		f.MethodByName("AddListener").Call(in)
	}
	return c
}

func (c *Config) DataChanged() {
	theme, _ := c.Theme.Get()
	fyne.CurrentApp().Preferences().SetString("theme", theme)

	language, _ := c.Language.Get()
	fyne.CurrentApp().Preferences().SetString("language", language)
}
