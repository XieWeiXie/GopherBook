package Apexcharts

type Chart struct {
	Height int    `json:"height"`
	Type   string `json:"type"`
}
type DataLabels struct {
	Enabled bool `json:"enabled"`
}
type Colors struct {
	Data []string `json:"colors"`
}

func (C *Colors) SetColors(color []string) {
	C.Data = color
}

type Series struct {
	Data []interface{} `json:"series"`
}

type XAxis struct {
	TypeName   string      `json:"type"`
	Categories interface{} `json:"categories"`
}

func (X *XAxis) SetCategories(data []string) {
	X.TypeName = "category"
	X.Categories = data
}

type Title struct {
	Text  string `json:"text"`
	Align string `json:"align"`
}

type Grid struct {
	Padding struct {
		Left   int `json:"left"`
		Right  int `json:"right"`
		Bottom int `json:"bottom"`
		Top    int `json:"top"`
	}
}
