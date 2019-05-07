package chapter4

import (
	"fmt"
	"sort"
)

func SortIntsUsage() {
	list := []int{10, 9, 2, 8, 3}
	sort.Ints(list)
	fmt.Println(list)
	sort.Sort(sort.Reverse(sort.IntSlice(list)))
	fmt.Println(list)
}

func SortFloatsUsage() {
	list := []float64{10, 9, 1.2, 3.4, 12.1}
	sort.Float64s(list)
	fmt.Println(list)
	sort.Sort(sort.Reverse(sort.Float64Slice(list)))
	fmt.Println(list)
}

func SortStringsUsage() {
	list := []string{"a", "A", "c", "C", "B", "b"}
	sort.Strings(list)
	fmt.Println(list)
	sort.Sort(sort.Reverse(sort.StringSlice(list)))
	fmt.Println(list)
}

type Language struct {
	Year    int    `json:"year"`
	Name    string `json:"name"`
	Account string `json:"account"`
}

type Languages []Language

func (ls Languages) Len() int {
	return len(ls)
}
func (ls Languages) Less(i, j int) bool {
	return ls[i].Year < ls[j].Year
}
func (ls Languages) Swap(i, j int) {
	ls[i], ls[j] = ls[j], ls[i]
}

func SortStruct() {
	list := Languages{
		{
			10, "Golang", "Google",
		}, {
			28, "Python", "Google",
		}, {
			30, "Java", "***",
		}, {
			1, "Lua", "23",
		},
	}
	fmt.Println("Ori", list)
	sort.Sort(list)
	fmt.Println("sort", list)
}
