package array

import "fmt"

type Array struct {
	data   []int
	length int
}

func NewArray(cap int) Array {
	return Array{
		data:   make([]int, cap),
		length: 0,
	}
}

func (A Array) IsFull() bool {
	return A.length == len(A.data)
}

func (A Array) IsEmpty() bool {
	return A.length == 0
}

func (A *Array) Insert(index, number int) (bool, error) {
	if A.IsFull() {
		return false, fmt.Errorf("full array")
	}
	if index > len(A.data)-1 || index < 0 {
		return false, fmt.Errorf("index out of range")
	}
	A.data[index] = number
	A.length++
	return true, nil

}

func (A *Array) Delete(number int) (bool, error) {
	if A.IsEmpty() {
		return false, fmt.Errorf("array is nil")
	}
	var v int
	var found bool
	for index, i := range A.data {
		if i == number {
			v = index
			found = true
			break
		}
	}
	if !found {
		return false, fmt.Errorf("not found number")
	}

	for i := v; i < A.length-1; i++ {
		A.data[i] = A.data[i+1]
	}

	A.length--
	A.data[A.length] = 0
	return true, nil
}

func (A Array) Print() {
	if A.IsEmpty() {
		fmt.Println("empty")
		return
	}
	for _, i := range A.data {
		fmt.Println(i)
	}
}

func (A Array) Cap() int {
	return A.length
}
