package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		number        int     = 1
		numberFloat   float32 = 1.23
		stringExample string  = "hello world"
		char                  = `A`
		isChar        bool    = false

		numberList = [3]int{1, 2, 3}
		stringList = [...]string{"212", "234", "345"}

		numberSlice = numberList[1:]
		stringSlice = stringList[1:]
	)

	type Info struct {
		Name       string
		Age        int
		University string
		Habit      map[string]string
	}

	habits := make(map[string]string)
	habits["One"] = "Go"
	habits["Two"] = "Python"

	var info = Info{
		Name:       "xieWei",
		Age:        23,
		University: "ShangHai",
		Habit:      habits,
	}

	var helloWorld func()
	helloWorld = func() {
		fmt.Println("Hello world")
	}

	var noName interface{}
	noName = 1
	noName = "12"
	noName = 1.23

	helloWorld()
	fmt.Println(number,
		numberFloat,
		stringExample,
		char,
		isChar,
		numberList,
		stringList,
		numberSlice,
		stringSlice,
		info,
		noName)

	var price = 100
	fmt.Println(strconv.Itoa(price))

	var isPrice bool = false
	fmt.Println(strconv.FormatBool(isPrice))

	type XieWei int
	type Zhihu string
	type WeChat Zhihu

	var xieWeiName XieWei
	var xieWeiZhihu Zhihu
	var xieWeiWechat WeChat
	xieWeiName = 12
	xieWeiZhihu = "Learn golang"
	xieWeiWechat = "Step by Step"

	fmt.Println(xieWeiName, xieWeiZhihu, xieWeiWechat)

}
