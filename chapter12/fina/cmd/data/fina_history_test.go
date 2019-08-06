package data

import (
	"fmt"
	"os"
	"testing"
)

func TestParseFiFaHistoryByXpath(t *testing.T) {
	f, err := os.Open("fifa_history.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	ParseFiNaHistoryByXpath(f)
}
