package assistance

import (
	"fmt"
	"strings"
	"testing"
)

func TestChromedpGetContent(t *testing.T) {
	content := ChromedpGetContent("http://quotes.toscrape.com/js/")
	fmt.Println(strings.Contains(content, `<div class="quote">`))
}
