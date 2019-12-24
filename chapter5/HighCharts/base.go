package HighCharts

type Chart struct {
	TypeName    string    `json:"type"`
	MarginTop   int       `json:"marginTop"`
	MarginRight int       `json:"marginRight"`
	Options3d   Options3d `json:"options3d"`
}

type Title struct {
	Text string `json:"text"`
}
type Series struct {
	Data []OneSeries `json:"series"`
}
type OneSeries struct {
	Name  string      `json:"name"`
	Data  interface{} `json:"data"`
	Stack string      `json:"stack"`
}

type XAxis struct {
	Categories []string `json:"categories"`
}

func (A *XAxis) AddCategories(data []string) {
	A.Categories = data
}

type YAxis struct {
	AllowDecimals bool  `json:"allowDecimals"`
	Min           int   `json:"min"`
	Title         Title `json:"title"`
}

func (Y *YAxis) AddYAxisTitle(title string) {
	Y.Title = Title{
		Text: title,
	}
}
func (Y *YAxis) AllowDecimal() {
	Y.AllowDecimals = true
}
