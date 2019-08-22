package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	array := NewArray(3)
	fmt.Println(array.IsEmpty(), array.IsFull(), array.Cap())
	array.Insert(2, 10)
	array.Insert(0, 100)
	array.Insert(1, 1000)

	array.Print()
	fmt.Println(array.Cap())
	isDeleted, _ := array.Delete(1000)
	fmt.Println(isDeleted, array.Cap())
	array.Print()
}
