package quotes

import "testing"

func TestGetQuotesContent(t *testing.T) {
	GetQuotesContent("http://quotes.toscrape.com/js/")
}

func TestGetQuotesContentByClick(t *testing.T) {
	GetQuotesContentByClick("http://quotes.toscrape.com/js/")
}
