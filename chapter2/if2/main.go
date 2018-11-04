package main

import "fmt"

func HelloGolang(language string) {

	if language == "Go" {
		fmt.Println("Hello " + language)
	}

}

func HelloGo(language string) {
	if language == "Go" {
		fmt.Println("Hello " + language)
	} else {
		fmt.Println("Hello ? " + language)
	}
}

func HelloLanguage(language string) {

	if language == "Go" {
		fmt.Println("Go Google")
	} else if language == "Python" {
		fmt.Println("import this")
	} else {
		fmt.Println("？？ " + language)
	}
}

func SayGo(number int) {

	for i := 0; i < number; i++ {
		fmt.Println("Hello Golang")
	}
}

func useSwitch(str string) {
	switch str {
	case "A":
		fmt.Println("Score >=", 90)
	case "B":
		fmt.Println("80 <= Score < 90")
	case "C":
		fmt.Println("70 <= Score < 80")
	case "D":
		fmt.Println("60 <= Score < 70")
	case "E":
		fmt.Println("50 <= Score < 60")
	default:
		fmt.Println("Score < 50")
	}
}

func useForRange(names []string) {
	for index, name := range names {
		fmt.Println(index, name)
	}
}

func main() {
	HelloGolang("Go")

	HelloGo("Go")
	HelloGo("Python")

	HelloLanguage("Go")
	HelloLanguage("Python")
	HelloLanguage("Java")

	SayGo(3)

	useSwitch("A")
	useSwitch("B")
	useSwitch("D")
	useSwitch("dad")

	useForRange([]string{"Zhao", "Qian", "Sun", "Li", "Zhou", "Wang"})

}
