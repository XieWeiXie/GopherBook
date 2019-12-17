package data

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/configs"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/assistance"
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
