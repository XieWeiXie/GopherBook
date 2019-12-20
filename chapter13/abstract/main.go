package main

import (
	"fmt"
	"log"
	"net/http"
)

type Gopher interface {
	Go(string) string
}
type Pythoneer interface {
	Python(string) string
}

type JobGo struct {
}

func (J JobGo) Go(v string) string {
	return fmt.Sprintf("实现了 Gopher Interface")
}

func (J JobGo) Python(v string) string {
	return fmt.Sprintf("同样实现了 Pythoneer Interface")
}

type JobPython struct {
}

func (J JobPython) Python(v string) string {
	return fmt.Sprintf("实现了 Pythoneer Interface")
}

type AwesomeDeveloper struct {
	JobGo
	JobPython
}

func (A AwesomeDeveloper) Go(v string) string {
	return fmt.Sprintf("实现了面向对象的多态特性")
}
func (A AwesomeDeveloper) Python(v string) string {
	return fmt.Sprintf("实现了面向对象的多态特性")
}

func ExampleWithGopher(v Gopher, value string) {
	log.Println(v.Go(value))
}

func ExampleWithPythoneer(v Pythoneer, value string) {
	log.Println(v.Python(value))
}

func main() {
	var g JobGo
	ExampleWithGopher(g, "JobGo")
	ExampleWithPythoneer(g, "JobGo")

	var p JobPython
	ExampleWithPythoneer(p, "JobPython")

	var awesome AwesomeDeveloper
	ExampleWithGopher(awesome, "AwesomeDeveloper")
	ExampleWithPythoneer(awesome, "AwesomeDeveloper")
}

type Parser interface {
	Fetch(url string) *http.Response
	Clean(response *http.Response) *http.Response
	IntoDB(client interface{}) bool
}
