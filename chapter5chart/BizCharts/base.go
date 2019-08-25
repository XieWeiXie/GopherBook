package BizCharts

import (
	"fmt"
	"net/http"
)

type BizChartInterface interface {
	Plot(w http.ResponseWriter, r *http.Request)
	Save(string) bool
	Name() string
	Type() string
}

type BaseData struct {
	Data interface{} `json:"data"`
}

type OneData map[string]interface{}

type BaseCols map[string]OneCol

func (B BaseCols) Keys() []string {
	var keys []string
	for k, _ := range B {
		keys = append(keys, k)
	}
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
