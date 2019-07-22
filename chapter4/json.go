package chapter4

import (
	"encoding/json"
	"fmt"
)

type JsonExample struct {
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age"`
	School string `json:"university"`
}

func JsonMarshal() {
	var jex JsonExample
	jex = JsonExample{
		Name:   "Go",
		Age:    10,
		School: "Google",
	}
	by, _ := json.Marshal(jex)
	fmt.Println(string(by))

}

func JsonUnmarshal() {

	var v JsonExample

	by := []byte(`{"name":"Go","age":10, "university":"google"}`)

	json.Unmarshal(by, &v)
	fmt.Println(v)

	var vother JsonExample
	byOther := []byte(`{"name":"","age":10, "school":"google"}`)
	json.Unmarshal(byOther, &vother)
	fmt.Println(vother)
}
