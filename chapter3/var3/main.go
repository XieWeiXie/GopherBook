package main

import (
	"fmt"
	"reflect"
)

// 显式的声明
func varDeclare() {

	var number int
	var name string

	number = 100
	name = "XieWei"

	fmt.Println(number, name)

}

// 隐式的声明
func varDeclareHide() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func manyVarDeclare() {

	var numberOne, numberTwo, numberThree int
	var name string
	numberOne, numberTwo, numberThree, name = 1, 2, 3, "XieWei"

	fmt.Println(numberOne, numberTwo, numberThree, name)
}

func manyVarDeclareBlock() {

	var (
		numberOne, numberTwo, numberThree int
		name                              string
	)
	numberOne, numberTwo, numberThree = 1, 2, 3
	name = "XieWei"

	fmt.Println(numberOne, numberTwo, numberThree, name)

}

func fetchNumberListMax(values []int) int {

	if len(values) < 1 {
		return 0
	}
	if len(values) == 1 {
		return values[0]
	}
	var numberMax int
	numberMax = values[0]

	for _, value := range values {
		if numberMax < value {
			numberMax = value
		}
	}
	return numberMax

}

func fetchNumberListMin(values []int) int {

	if len(values) < 1 {
		return 0
	}
	if len(values) == 1 {
		return values[0]
	}
	var numberMin int
	numberMin = values[0]

	for _, value := range values {

		if numberMin > value {
			numberMin = value
		}
	}
	return numberMin

}

type Info struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Number int    `json:"number"`
}

func hasInfoName(value string) bool {

	var info Info
	info.Name = "XieWei"
	info.Age = 20
	info.Number = 100

	var typeInfo reflect.Type
	typeInfo = reflect.TypeOf(info)

	if _, ok := typeInfo.FieldByName(value); ok {
		return ok
	}
	return false

}

func main() {
	varDeclare()
	varDeclareHide()

	manyVarDeclare()
	manyVarDeclareBlock()

	var values = []int{1, 2, 3, 4, 5, 200, 6, 7, 8, 9, 100}
	min := fetchNumberListMin(values)
	max := fetchNumberListMax(values)
	fmt.Println(min, max)

	fmt.Println(hasInfoName("Name"))
}
