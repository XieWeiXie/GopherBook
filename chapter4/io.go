package chapter4

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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

func ReadFile() {
	f, err := os.Open("io_test.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
	}
}
