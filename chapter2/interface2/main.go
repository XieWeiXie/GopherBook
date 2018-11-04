package main

import "fmt"

type Controller interface {
	SayHello()
	SayNumber(int)
	SayHi()
}

type DefaultController struct {
}

func (d DefaultController) SayHello() {
	fmt.Println("Hello world")
}

func (d DefaultController) SayNumber(number int) {
	fmt.Println(fmt.Sprintf("%d", number))
}

func (d DefaultController) SayHi() {
	fmt.Println("Say Hi")
}

type ErrorCode struct {
	Code    int
	Message string
}

func (e ErrorCode) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func SayError() error {
	var e ErrorCode
	e.Code = 400
	e.Message = "http status code"
	return e
}

func main() {

	var d DefaultController
	var c Controller
	c = d

	c.SayHello()
	c.SayNumber(123)
	c.SayHi()

	SayError()

}
