package assistance

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestGetContent(t *testing.T) {
	content, err := GetContent("http://www.baidu.com/")
	if err != nil {
		log.Println(err)
		return
	}
	//fmt.Println(string(content))
	fmt.Println(strings.Contains(string(content), "百度一下，你就知道"))
}

func TestGetResponse(t *testing.T) {
	content, err := GetResponse("http://www.baidu.com/")
	if err != nil {
		log.Println(err)
		return
	}
	//fmt.Println(string(content))
	fmt.Println(strings.Contains(string(content), "百度一下，你就知道"))
}
