package data

import (
	"fmt"
	"testing"
)

func TestGetDate(tests *testing.T) {
	v := "2009-12-18T00:00:00"
	fmt.Println(getDate(v))
}
