package chapter4

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func IoUsage() {
	url := "http://httpbin.org/anything?name=xix"
	request, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(`{"name":"XieXie"}`))
	client := http.DefaultClient
	response, _ := client.Do(request)
	defer response.Body.Close()
	by, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(by))
}
