package Echarts

import "strings"

type TitleOpts struct {
	Text      string `json:"text,omitempty"`
	TextAlign string `json:"textAlign,omitempty"`
	Top       string `json:"top,omitempty"`  //'top', 'middle', 'bottom'
	Left      string `json:"left,omitempty"` //'left', 'center', 'right'
	Right     string `json:"right,omitempty"`
	Bottom    string `json:"bottom,omitempty"`
}

const (
	AUTO = iota
	LEFT
	RIGHT
	CENTER
)

var DefaultTextAlign map[int]string

func init() {
	DefaultTextAlign = make(map[int]string)
	DefaultTextAlign[AUTO] = strings.ToLower("auto")
	DefaultTextAlign[LEFT] = strings.ToLower("left")
	DefaultTextAlign[RIGHT] = strings.ToLower("right")
	DefaultTextAlign[CENTER] = strings.ToLower("center")
}

const (
	BOTTOM = iota
	TOP
	MIDDLE
)

var DefaultTop map[int]string

func init() {
	DefaultTop = make(map[int]string)
	DefaultTop[BOTTOM] = strings.ToLower("bottom")
	DefaultTop[TOP] = strings.ToLower("top")
	DefaultTop[MIDDLE] = strings.ToLower("middle")
}

var DefaultLeft map[int]string

func init() {
	DefaultLeft = make(map[int]string)
	DefaultLeft[LEFT] = "left"
	DefaultLeft[RIGHT] = "right"
	DefaultLeft[CENTER] = "center"
}

func (T *TitleOpts) SetTextAlign(index int) {
	T.TextAlign = DefaultTextAlign[index]
}

func (T *TitleOpts) SetPositions(top int, left int) {
	T.Top = DefaultTop[top]
	T.Left = DefaultLeft[left]
}

type AxisOpts struct {
	Data interface{} `json:"data"`
}

type ToolTipOpts struct {
	Data interface{} `json:"data"`
}

type LegendOpts struct {
	Data interface{} `json:"data"`
}

type Series struct {
	Data []OneSeries `json:"series"`
}

type OneSeries struct {
	SymbolSize int         `json:"symbolSize,omitempty"`
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	Data       interface{} `json:"data"`
}

func (S *Series) Add(data ...OneSeries) {
	S.Data = append(S.Data, data...)
}
