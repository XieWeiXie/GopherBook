package chapter4

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func getHandle(rawString string) {
	response, err := http.Get(rawString)
	if err != nil {
		return
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func postHandle(rawString string, body io.Reader) {
	response, err := http.Post(rawString, "application/json", body)
	if err != nil {
		return
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func ClientUsage() {
	// get
	getHandle("http://localhost:80/headers")
	getHandle("http://localhost:80/ip")
	getHandle("http://localhost:80/user-agent")

	// post
	var buf bytes.Buffer
	buf.WriteString("hello golang")
	postHandle("http://localhost:80/anything", &buf)

	val := strings.NewReader("hello python")
	postHandle("http://localhost:80/anything", val)

	bytesNew := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "Golang",
		Age:  10,
	}
	byt, _ := json.Marshal(bytesNew)
	postHandle("http://localhost:80/anything", bytes.NewReader(byt))
	// PostForm
	response, err := http.PostForm("http://localhost:80/anything", url.Values{
		"name": []string{"Golang"},
		"age":  []string{"10"},
	})
	if err != nil {
		return
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	if strings.Contains(req.URL.Path, "header") {
		return errors.New("header")
	}
	return nil

}
func UserClientUsage() {
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:80/ip", nil)
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

type SelfHandler struct {
}

func (SelfHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("Hello Python"))
}

func ServerUsage() {
	// method One
	http.HandleFunc("/hello_golang", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello Golang"))
	})

	// method Two
	var self SelfHandler
	http.Handle("/hello_python", self)

	// method Three

	var selfServerMux *http.ServeMux
	selfServerMux = &http.ServeMux{}
	selfServerMux.HandleFunc("/hello_golang_2", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello Golang 2"))
	})

	var selfServer http.Server
	var selfHandler Self
	var selfMux *http.ServeMux
	selfMux = &http.ServeMux{}
	selfHandler = Self{}
	selfMux.Handle("/say", selfHandler)
	selfServer = http.Server{
		Handler: selfHandler,
		Addr:    "localhost:9099",
	}
	go func() {
		log.Fatal(http.ListenAndServe(":9090", selfServerMux))

	}()
	go func() {
		log.Fatal(selfServer.ListenAndServe())
	}()
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	select {}
}

type Self struct {
}

func (Self) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Hello Self Sever 1")
}

func (Self) Say(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Hello Self Sever 1")
}
