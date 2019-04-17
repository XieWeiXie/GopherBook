package chapter4

import (
	"fmt"
	"testing"
	"unicode"
)

func TestUnicode(test *testing.T) {
	UnicodeUsage()
	fmt.Println()
	fmt.Println(RegisterUserName("注册名Hello", unicode.Scripts["Han"]))
	fmt.Println(RegisterUserName("등록이름", unicode.Scripts["Hangul"]))

}
