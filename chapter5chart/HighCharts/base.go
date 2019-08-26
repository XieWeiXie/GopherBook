package HighCharts

type Chart struct {
	TypeName    string `json:"type"`
	MarginTop   int    `json:"marginTop"`
	MarginRight int    `json:"marginRight"`
	value       map[string]interface{}
}

func (C Chart) AddField(key string, v interface{}) {
	C.value[key] = v
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

type XAixs struct {
	Categories []string `json:"categories"`
}

func (A *XAixs) AddCategories(data []string) {
	A.Categories = data
}

type YAixs struct {
	AllowDecimals bool  `json:"allowDecimals"`
	Min           int   `json:"min"`
	Title         Title `json:"title"`
}

func (Y *YAixs) AddYAxisTitle(title string) {
	Y.Title = Title{
		Text: title,
	}
}
