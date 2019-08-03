package data

import (
	"GopherBook/chapter12/fina/configs"
	"GopherBook/chapter12/fina/pkg/assistance"
	"fmt"
	"testing"
)

func TestParseChampionshipByJquery(t *testing.T) {
	reader, err := assistance.DownloaderReturnIOReader(configs.MatchDescription)
	if err != nil {
		fmt.Println(err)
		return
	}
	a, _ := ParseChampionshipByJquery(reader)
	fmt.Println(a)
}
