package data

import (
	"GopherBook/chapter12/fina/configs"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestParseSportsByQuery(t *testing.T) {
	f, err := os.Open("sports.html")
	if err != nil {
		log.Println(err)
		return
	}

	ParseSportsByQuery(f)
}

func TestParseSportsByQuery2(t *testing.T) {
	for key, i := range configs.MatchSportsMap {
		url := fmt.Sprintf(configs.MatchSports, key, i)

	}

}
