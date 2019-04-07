package chapter4

import (
	"fmt"
	"strconv"
)

func ToBool() {
	var a bool
	a = true
	b := strconv.FormatBool(a)
	fmt.Println(b)
	c, _ := strconv.ParseBool("false")
	fmt.Println(c)

}

func ToNumber() {
	var (
		a int
		b uint64
		c float64
	)

	a = 1
	b = 2
	c = 3.14
	fmt.Println(strconv.Itoa(a))
	fmt.Println(strconv.FormatUint(b, 10))
	fmt.Println(strconv.FormatFloat(c, 'f', 1, 32))

	d := "4.178"
	floatD, _ := strconv.ParseFloat(d, 64)
	fmt.Println(floatD)
}
