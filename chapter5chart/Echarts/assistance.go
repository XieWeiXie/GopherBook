package Echarts

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func toHandler(w http.ResponseWriter, r *http.Request, data interface{}) {
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

func toSave(v interface{}, name string) bool {
	t, err := template.New("").Parse(PlotText())
	if err != nil {
		log.Println(err)
		return false
	}
	file, err := os.Open(name)
	if err != nil {
		log.Println(err)
		return false
	}
	err = t.Execute(file, v)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
