package BizCharts

import (
	"fmt"
)

type BaseTheme struct {
	Theme string `json:"theme"`
}

const (
	DEFAULT_THEME = "default"
	DARK_THEME    = "dark"
)

func (B *BaseTheme) SetTheme(theme string) {
	B.Theme = theme
}

type BaseData struct {
	Data interface{} `json:"data"`
}

type BaseCols struct {
	X OneCol
	Y OneCol
}

func (B BaseCols) Keys() []string {
	var keys []string
	keys = append(keys, B.X.Alias, B.Y.Alias)
	return keys
}

func (B BaseCols) Position() string {
	keys := B.Keys()
	if len(keys) != 2 {
		return "-1"
	}
	return fmt.Sprintf("%s*%s", keys[0], keys[1])
}

type OneCol struct {
	Alias string `json:"alias"`
}

type BaseLegend struct {
	Location string `json:"location"`
}
