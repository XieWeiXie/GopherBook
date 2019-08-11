package assistance

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateTime = time.Tick(time.Millisecond * 200)

func GetContent(url string) ([]byte, error) {
	<-rateTime
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
