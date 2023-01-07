package theme

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"reflect"
)

type Config struct {
	Theme binding.String
}

func NewConfig() *Config {
	c := &Config{
		Theme: binding.NewString(),
	}

	_ = c.Theme.Set(fyne.CurrentApp().Preferences().StringWithFallback("theme", "__DARK__"))

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
	// todo flush with field tag instead of the first param
	theme, _ := c.Theme.Get()
	fyne.CurrentApp().Preferences().SetString("theme", theme)
}
