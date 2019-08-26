package BizCharts

import "net/http"

type Polar struct {
	BaseData
	BaseLegend
	BaseTheme
	coordType    string
	geomType     string
	position     string
	color        string
	TemplateData TemplateData
}

type TemplateData struct {
	Data      interface{} `json:"data"`
	TypeCoord interface{} `json:"typeCoord"`
	TypeGeom  interface{} `json:"typeGeom"`
	Location  interface{} `json:"location"`
	Position  interface{} `json:"position"`
	Color     interface{} `json:"color"`
	Theme     interface{} `json:"theme"`
}

func (P *Polar) valueFormat() {
	P.TemplateData = TemplateData{
		Data:      P.Data,
		TypeCoord: P.coordType,
		TypeGeom:  P.geomType,
		Location:  P.Location,
		Position:  P.position,
		Color:     P.color,
		Theme:     P.Theme,
	}

}

func (P Polar) Plot(w http.ResponseWriter, r *http.Request) {
	P.valueFormat()
	toHandler(w, r, P.TemplateData)
}

func (P Polar) Name() string {
	return P.coordType
}
func (P Polar) String() string {
	return P.Name()
}
func (P Polar) Save(name string) bool {
	if name == "" {
		name = P.coordType
	}
	return toSave(name, P.TemplateData)
}

func NewPolar(theme string) *Polar {
	return &Polar{
		BaseLegend: BaseLegend{
			Location: "right",
		},
		BaseTheme: BaseTheme{
			Theme: theme,
		},
		geomType:  "interval",
		coordType: POLARTYPE,
	}
}
