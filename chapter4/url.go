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

func UrlUsageParams() {
	rawUrl := "https://golang.org/pkg/net/url?name=xie&age=20"
	urlParsed, _ := url.Parse(rawUrl)
	fmt.Println(urlParsed.Query())
	v := urlParsed.Query()
	v.Del("name")
	v.Add("school", "shanghai")
	urlParsed.RawQuery = v.Encode()
	fmt.Println(urlParsed)
}

func UrlValues() {
	values := "name=xie&age=20"
	v, _ := url.ParseQuery(values)
	fmt.Println(v)
	v.Add("school", "shanghai")
	fmt.Println(v)
}
