package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

var ErrExampleNew = errors.New("hello world error")
var ErrExampleFmt = fmt.Errorf("hello world %s", "error")

type ErrorMessage struct {
	Err     error
	Code    int
	Message string
}

func (e *ErrorMessage) Error() string {
	return fmt.Sprintf("e.err = %s, e.code = %d, e.message = %s", e.Err.Error(), e.Code, e.Message)
}

var (
	ErrNotRoute   = ErrorMessage{Err: errors.New("not route"), Code: 404, Message: "check route"}
	ErrParamNotOk = ErrorMessage{Err: errors.New("param not ok"), Code: 10000, Message: "check param"}
)

type University struct {
	Name      string  `json:"name"`
	Location  string  `json:"location"`
	Number    float64 `json:"student_number,string"`
	President string  `json:"-"`
}

func (u University) MarshalJSON() ([]byte, error) {

	result := fmt.Sprintf("name: %% %s, location: %% %s , 人数: %% %f", u.Name, u.Location, u.Number)
	return json.Marshal(result)
}

func main() {

	var university University
	university = University{
		Name:      "ShangHaiUniversity",
		Location:  "ShangHai",
		Number:    2000000,
		President: "XXXXX",
	}

	universityByte, _ := json.Marshal(university)
	fmt.Println(string(universityByte))
}
