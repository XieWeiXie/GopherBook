package assistance

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestSeleniumGetContent(t *testing.T) {
	contentOne, err := SeleniumGetContent("http://quotes.toscrape.com/js/")
	if err != nil {
		log.Println(err)
		return
	}
	contentTwo, err := GetContent("http://quotes.toscrape.com/js/")
	fmt.Println(strings.Contains(contentOne, `<div class="quote">`))
	fmt.Println(strings.Contains(string(contentTwo), `<div class="quote">`))

}
