package chapter4

import (
	"errors"
	"fmt"
	"unicode"
)

func UnicodeUsage() {

	var string = "你好 Golang 123"

	for _, i := range string {
		if unicode.IsLetter(i) {
			fmt.Printf("Yes:%c ", i)
		} else {
			fmt.Printf("No:%c ", i)
		}

	}
	fmt.Println()
	for _, i := range string {
		if unicode.Is(unicode.Scripts["Han"], i) {
			fmt.Printf("%c\n", i)
		}
	}

	for _, i := range string {
		fmt.Printf("%c", unicode.ToUpper(i))
	}

}

func RegisterUserName(name string, table *unicode.RangeTable) error {
	for _, i := range name {
		if !unicode.Is(table, i) {
			return errors.New("scripts is not correct")
		}
	}
	return nil
}
