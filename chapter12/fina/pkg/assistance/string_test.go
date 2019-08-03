package assistance

import (
	"fmt"
	"testing"
)

func TestGetWords(t *testing.T) {
	tests := []struct {
		value string
	}{
		{"Advance向着未来梦想全力迈入"},
		{"Environment感受自然气息的环境"},
	}
	for _, i := range tests {
		fmt.Println(GetWordsExceptHan(i.value))
		fmt.Println(GetWordsHan(i.value))
	}
}

func TestGetDate(t *testing.T) {
	tests := []struct {
		val string
	}{
		{
			val: "2019年07月12日-07月28日(17天)",
		},
	}
	for _, i := range tests {
		fmt.Println(GetDate(i.val))
	}
}
