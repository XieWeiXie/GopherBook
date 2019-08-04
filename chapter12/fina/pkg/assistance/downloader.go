package assistance

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var rateTime = time.Tick(2000 * time.Millisecond)

func requestSet(url string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36")
	request.Header.Add("Origin", "https://www.fina-gwangju2019.com")
	request.Header.Add("Host", "www.fina-gwangju2019.com")
	return request
}

func Downloader(url string) ([]byte, error) {
	<-rateTime
	request := requestSet(url)
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	bodyReader := bufio.NewReader(response.Body)
	e := DetermineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error %v\n", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func DownloaderReturnIOReader(url string) (io.Reader, error) {
	<-rateTime
	request := requestSet(url)
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return charset.NewReader(response.Body, response.Header.Get("Content-type"))
}

func PostReturnIOReader(router string, body io.Reader) ([]byte, error) {
	<-rateTime
	request, err := http.NewRequest(http.MethodPost, router, body)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	content, err := charset.NewReader(response.Body, response.Header.Get("Content-type"))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(content)

}
