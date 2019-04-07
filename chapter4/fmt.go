package chapter4

import (
	"fmt"
)

func FmtUsage() {
	var number = 100.203
	var numberInt = 100
	fmt.Printf("%d\n", numberInt)
	fmt.Printf("%o\n", numberInt)
	fmt.Printf("%x\n", numberInt)
	fmt.Printf("%X\n", numberInt)
	fmt.Printf("%b\n", numberInt)
	fmt.Printf("%f\n", number)
	fmt.Printf("%e\n", number)
	fmt.Printf("%E\n", number)
}

func FmtStringUsage() {
	var values = "golang"
	fmt.Printf("%s\n", values)
	fmt.Printf("%q\n", values)
}

func FmtBoolUsage() {
	var ok = true
	fmt.Printf("%t\n", ok)
}

func FmtOtherUsage() {
	var a = 1
	var b = 2.0
	var ok = true
	number := &a
	var s = struct {
		Name string `json:"name"`
	}{
		Name: "Go",
	}
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", ok)
	fmt.Printf("%p\n%d\n", &a, number)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)
}

type Val struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (v Val) String() string {
	return fmt.Sprintf("%s + %d", v.Name, v.Age)
}
func (v Val) GoString() string {
	return fmt.Sprintf("%s + %d", v.Name, v.Age)
}
