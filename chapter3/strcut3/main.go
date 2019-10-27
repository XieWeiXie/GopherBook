package main

import (
	"fmt"
	"unsafe"
)

type Info struct {
	Name string
	Age  int
}

type Student struct {
	Name string
	University
}

func (s Student) PrintName() {
	fmt.Println(s.Name)
}

type University struct {
	Name     string
	Location string
}

func (u University) PrintName() {
	fmt.Println(u.Name)
}

func main() {
	var infoOne Info = Info{
		Name: "XieWei",
		Age:  20,
	}

	var inofTwo = Info{"XieWei", 20}

	var infoThree = new(Info)

	infoThree = &Info{
		Name: "XieWei",
		Age:  20,
	}

	fmt.Println("One", infoOne)
	fmt.Println("Two", inofTwo)
	fmt.Println("Three", *infoThree)
	fmt.Println(unsafe.Sizeof(infoOne), fmt.Sprintf("%x - %d - %x - %d", &infoOne.Name, unsafe.Sizeof(infoOne.Name), &infoOne.Age, unsafe.Sizeof(infoOne.Age)))
	fmt.Println(unsafe.Sizeof(inofTwo), fmt.Sprintf("%x - %d - %x - %d", &inofTwo.Name, unsafe.Sizeof(inofTwo.Name), &inofTwo.Age, unsafe.Sizeof(inofTwo.Age)))
	fmt.Println(unsafe.Sizeof(*infoThree), fmt.Sprintf("%x - %d - %x - %d", &infoThree.Name, unsafe.Sizeof(infoThree.Name), &infoThree.Age, unsafe.Sizeof(infoThree.Age)))

	var std Student
	std.Name = "XieWei"
	std.University.Name = "ShangHai"
	std.Location = "ShangHai"
	fmt.Println(std)

	std.PrintName()
	std.University.PrintName()

}
