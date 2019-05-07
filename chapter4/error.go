package chapter4

import (
	"errors"
	"fmt"
)

func ErrorUsage() {

	err := errors.New("err: found 1")
	if err != nil {
		fmt.Println(err.Error())
	}
	err2 := fmt.Errorf("err: %s", "found 2")
	if err2 != nil {
		fmt.Println(err2.Error())
	}
}

type SelfError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (self SelfError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", self.Code, self.Message)
}

func UsageError(value string) error {
	var self SelfError
	if value == "" {
		self.Code = 400
		self.Message = "fail"
		return self
	}
	return nil
}

func UserErrorUsage() {

	err3 := UsageError("")
	if err3 != nil {
		fmt.Println(err3.Error())
	}

}
