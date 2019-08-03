package data

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestParseFiFaBriefByjQuery(t *testing.T) {
	f, err := os.Open("fifa_brief.html")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(ParseFiFaBriefByjQuery(f))
}
