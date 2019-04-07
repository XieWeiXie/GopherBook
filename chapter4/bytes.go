package chapter4

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ToString(value []byte) string {
	return string(value)
}

func ToBytes(value string) []byte {
	return []byte(value)
}

func HttpByBytes() {
	url := "http://httpbin.org/anything?name=xix"

	var body map[string]string
	body = make(map[string]string)
	body["age"] = "20"
	body["school"] = "ShangHai"

	by, _ := json.Marshal(body)

	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(by))
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func HttpByByteNewReader() {
	url := "http://httpbin.org/anything?name=xix"

	var body map[string]string
	body = make(map[string]string)
	body["age"] = "20"
	body["school"] = "ShangHai"

	by, _ := json.Marshal(body)

	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(by))
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func HttpByStrings() {
	url := "http://httpbin.org/anything?name=xix"

	request, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(`{"name":"XieWei", "school":"ShangHai"}`))
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
