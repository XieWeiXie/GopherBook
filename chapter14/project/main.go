package main

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/chapter14/project/internal"
)

func main() {
	v := internal.NewVersion("v0.14")
	fmt.Println(v.GetValue())
}
