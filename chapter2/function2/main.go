package main

import (
	"fmt"
	"strconv"
)

var PrintName = func() {
	fmt.Println("Hello Name")
}

func PrintHello() {
	fmt.Println("Hello world")
}

func SumNumber(numberOne int, numberTWo int) int {
	return numberOne + numberTWo
}

func NameResult(numberOne, numberTwo int) (result int) {
	result = numberOne + numberTwo
	return
}

func MultiResult(numberOne int, numberTwo int) (int, string) {
	sum := numberOne + numberTwo
	return sum, strconv.Itoa(sum)
}

var anonymousFuncTimes = func(numberOne int) int {
	return numberOne * 10

}

func main() {
	PrintHello()
	PrintName()
	sumAdd := SumNumber(1, 2)
	fmt.Println(sumAdd)
	sumResult := NameResult(10, 20)
	fmt.Println(sumResult)

	fmt.Println(MultiResult(100, 200))

	fmt.Println(anonymousFuncTimes(100))
}
