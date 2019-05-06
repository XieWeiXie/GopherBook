package chapter4

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ClientUsage() {
	// get
	response, err := http.Get("http://httpbin.org/get?name=xie")
	if err != nil {
		return
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	// post
	//request, _ := http.NewRequest(http.MethodPost)
}

func ServerUsage() {}
