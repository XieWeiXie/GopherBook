package main

import "fmt"

func main() {

	// 变量
	var number int
	number = 1

	numberTwo := 2

	var stringExample string
	stringExample = "How are you"

	var numberList []int
	numberList = []int{1, 2, 3, 4}

	var isNumber bool
	isNumber = true

	canFly := false

	fmt.Println(number, numberTwo, stringExample, numberList, isNumber, canFly)

	// 多个变量声明和赋值
	var strOne, strTwo, strThree string
	strOne, strTwo, strThree = "1", "2", "3"

	numOne, numTwo, numThree := 1, 2, 4

	fmt.Println(strOne, strTwo, strThree, numOne, numTwo, numThree)

	// 作用域操作
	varFunc()
	fmt.Println(exampleNumber, exampleString)
	varGlobalFunc()

	// 常量
	const Name string = "xieWei"
	const Age int = 25
	const Info string = "ShangHai"
	fmt.Println(Name, Age, Info)
}

var exampleNumber int = 1156143589
var exampleString string = "WuXiaoShen"

func varFunc() {

	var exampleNumber = 987654321

	fmt.Println(exampleNumber)
}

func varGlobalFunc() {
	fmt.Println(exampleNumber)
}
