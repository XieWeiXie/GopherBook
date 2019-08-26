package Apexcharts

import (
	"log"
	"net/http"
	"os"
	"text/template"
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

func toSave(name string, data interface{}) bool {
	f, err := os.Create(name)
	if err != nil {
		log.Println(err)
		return false
	}
	t, err := template.New("").Parse(PlotText())
	if err != nil {
		log.Println(err)
		return false
	}
	err = t.Execute(f, data)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
