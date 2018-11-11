package main

import (
	"fmt"
	"reflect"
)

var add = func(numberOne int, numberTwo int) int {
	return numberOne + numberTwo
}

var minus = func(numberOne int, numberTwo int) int {
	return numberOne - numberTwo
}

var multiply = func(number int, price float64) float64 {
	return float64(number) * price
}

var division = func(numberOne float64, numberTwo int) float64 {
	return float64(float64(numberOne) / float64(numberTwo))
}

var judgeMarry = func(manAge int, womanAge int) bool {
	if manAge >= 22 && womanAge >= 20 {
		return true
	}
	return false
}

var opList = func(number [4]int) {
	fmt.Println(number[1], reflect.TypeOf(number[1]))
	fmt.Println(len(number))
	fmt.Println(number[1:], reflect.TypeOf(number[1:]))

	for index, one := range number {
		fmt.Println(index, one)
	}

	for i := 0; i < len(number); i++ {
		fmt.Println(i, number[i])
	}
}

var opSlice = func(name []string) []string {
	fmt.Println(name[1], reflect.TypeOf(name[1]))

	for index, one := range name {
		fmt.Println(index, one)
	}

	name = append(name, "XieWei")

	return name
}

var opMap = func(name map[string]int) map[string]int {

	for key, value := range name {
		fmt.Println(key, value)
	}

	name["Life"] = 100

	if value, ok := name["Go"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("no exists Go")

	}
	delete(name, "java")

	return name
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(minus(1, 2))
	fmt.Println(multiply(1, 2))
	fmt.Println(division(1.234, 2))

	fmt.Println(judgeMarry(20, 20))
	fmt.Println(judgeMarry(25, 26))
	fmt.Println(judgeMarry(18, 20))

	var number [4]int = [...]int{1, 2, 3, 4}
	opList(number)

	var name []string = []string{"Go", "Python", "Java", "C++", "C#"}
	fmt.Println(opSlice(name))

	fmt.Println(make([]int, 2))
	fmt.Println(new([]int))

	nameMap := make(map[string]int)
	nameMap["java"] = 200
	nameMap["php"] = 100
	nameMap["python"] = 180
	nameMap["js"] = 220

	fmt.Println(opMap(nameMap))

}
