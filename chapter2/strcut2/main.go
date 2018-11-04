package main

import "fmt"

type Info struct {
	Name       string
	Age        int
	University string
}

type MyInt struct {
	Number int
}

func (m MyInt) SayHello() {
	fmt.Println("Hello World")
}

func (m *MyInt) SetNumber(other int) {
	m.Number = other
}

func (m MyInt) SayNumber() {
	fmt.Println(m.Number)
}

type ViewName struct {
	Name string
	ViewOther
}

func (v ViewName) SayName() {
	fmt.Println(v.Name)
}

type ViewOther struct {
	Value string
}

func (v ViewOther) SayValue() {
	fmt.Println(v.Value)
}

func main() {

	var info Info
	info = Info{
		Name:       "XieWei",
		Age:        20,
		University: "ShangHai",
	}

	var infoTwo = new(Info)
	infoTwo.Name = "XieWei"
	infoTwo.Age = 22
	infoTwo.University = "BeiJing"

	fmt.Println(info, infoTwo, *infoTwo)

	var my MyInt
	my.Number = 1
	my.SayHello()
	my.SayNumber()
	my.SetNumber(100)
	my.SayNumber()

	var viewName ViewName
	viewName.Name = "xieWei"
	viewName.Value = "value"
	viewName.SayName()
	viewName.SayValue()

}
