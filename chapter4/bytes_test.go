package chapter4

import (
	"fmt"
	"testing"
)

func TestConvert(tests *testing.T) {

	var A []byte
	A = []byte("a b")

	var B string
	B = "a b"

	fmt.Println(fmt.Sprintf("%T,%T", A, B))
	fmt.Println(fmt.Sprintf("%T,%T", ToString(A), ToBytes(B)))

	fmt.Println("1")
	HttpByBytes()
	fmt.Println("2")
	HttpByByteNewReader()
	fmt.Println("3")
	HttpByStrings()
}
