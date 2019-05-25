package main

import (
	"GopherBook/Chapter5/optimization/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	r := router.New()
	r.GET("/get", func(writer http.ResponseWriter, request *http.Request, ps router.Params) {
		writer.Write([]byte(fmt.Sprintf("method: %s, url: %s", request.Method, request.RequestURI)))
	})
	r.POST("/post", func(writer http.ResponseWriter, request *http.Request, ps router.Params) {
		writer.Write([]byte(fmt.Sprintf("method: %s, url: %s", request.Method, request.RequestURI)))
	})
	r.PATCH("/patch", func(writer http.ResponseWriter, request *http.Request, ps router.Params) {
		writer.Write([]byte(fmt.Sprintf("method: %s, url: %s", request.Method, request.RequestURI)))
	})
	r.DELETE("/delete", func(writer http.ResponseWriter, request *http.Request, ps router.Params) {
		writer.Write([]byte(fmt.Sprintf("method: %s, url: %s", request.Method, request.RequestURI)))
	})

	log.Fatal(http.ListenAndServe(":9090", r))

}
