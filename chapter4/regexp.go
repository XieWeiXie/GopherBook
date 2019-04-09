package chapter4

import (
	"fmt"
	"regexp"
)

func SimpleUsage() {
	Slogan := "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."

	reg, _ := regexp.Compile("open source programming language")
	if reg.Match([]byte(Slogan)) {
		fmt.Println("byte: Match")
	}

	if reg.MatchString(Slogan) {
		fmt.Println("string: Match")
	}

}

func SearchUsage() {
	Slogan := "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."

	reg, _ := regexp.Compile("open source programming language")
	v := reg.Find([]byte(Slogan))
	fmt.Println(string(v))

	v2 := reg.FindString(Slogan)
	fmt.Println(v2)
}

func ReplaceUsage() {
	Slogan := "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."

	reg, _ := regexp.Compile(`^Go`)
	result := reg.ReplaceAllString(Slogan, "Python")

	fmt.Println(result)
}

func SplitUsage() {
	Slogan := "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."

	reg, _ := regexp.Compile(`\s|\,|\.`) // 按" "、","、"." 分割

	result := reg.Split(Slogan, -1)
	fmt.Println(result)
}
