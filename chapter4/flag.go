package chapter4

import (
	"flag"
	"fmt"
)

func FlagUsage() {
	var number int
	flag.IntVar(&number, "n", 10, "number")

	name := flag.String("name", "Go", "name of language")

	flag.Parse()

	flag.PrintDefaults()
	fmt.Println(number, *name)
}
