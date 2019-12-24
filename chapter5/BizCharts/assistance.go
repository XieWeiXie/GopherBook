package BizCharts

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

func toHandler(w http.ResponseWriter, _ *http.Request, data interface{}) {
	t, err := template.New("").Parse(PlotText())
	if err != nil {
		log.Println(err)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		log.Println(err)
		return
	}
}

func toSave(name string, data interface{}) bool {
	f, err := os.Create(name)
	if err != nil {
		log.Println(err)
		return false
	}
	t, _ := template.New("").Parse(PlotText())
	err = t.Execute(f, data)
	if err != nil {
		log.Println(err)
		return false
	}
	return true

}
