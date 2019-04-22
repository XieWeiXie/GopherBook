package chapter4

import (
	"fmt"
	"net/url"
)

func UrlUsage() {
	var urlString = "https://golang.org/pkg/net/url?name=xie&age=20"
	urlPath, _ := url.Parse(urlString)
	fmt.Println(fmt.Sprintf("%#v", urlPath))
	v := urlPath.Query()
	v.Set("name", "Wei")
	urlPath.RawQuery = v.Encode()
	fmt.Println(fmt.Sprintf("%#v", urlPath))
}
