package chapter4

import (
	"fmt"
	"strings"
	"testing"
)

func TestAllFunction(tests *testing.T) {
	// 1
	fmt.Println(StringsContains("Go"))
	fmt.Println(StringsContains("Java"))

	// 2
	fmt.Println(StringsCompare("Java", "Go"))
	fmt.Println(StringsCompare("Go", "Java"))
	fmt.Println(StringsCompare("A", "B"), rune('A'), rune('B'))

	// 3
	fmt.Println(StringsToUpper("goLang, hello world"))
	fmt.Println(StringsToLower("GoLang"))
	fmt.Println(StringsToTitle("goLang, hello world"))

	// 4
	fmt.Println(StringsCount("Go"))
	fmt.Println(StringsCount("s"))

	// 5
	fmt.Println(StringsHasSuffix("software"))
	fmt.Println(StringsHasSuffix("software."))
	fmt.Println(StringsHasPrefix("Java"))
	fmt.Println(StringsHasPrefix("Go"))

	// 6
	fmt.Println(StringsSplit(","), len(StringsSplit(",")))
	fmt.Println(StringsJoin([]string{"Go", "Java", "Python"}))

	// 7
	fmt.Println(StringsIndex("o"))

	// 8
	fmt.Println(StringsTrim("   hello world   "))

	// 9
	fmt.Println(StringsReplacer(" hello world ,\n golang"))
	var a strings.Builder
	a.WriteString("ad")
	fmt.Println(a.String())
}
