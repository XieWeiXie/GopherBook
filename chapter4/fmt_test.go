package chapter4

import (
	"fmt"
	"testing"
)

func TestFmtUsage(tests *testing.T) {
	FmtUsage()
	FmtStringUsage()
	FmtBoolUsage()
	FmtOtherUsage()

	var a = Val{
		Name: "go",
		Age:  20,
	}
	fmt.Println(a)
	fmt.Printf("%#v\n", a)
}
