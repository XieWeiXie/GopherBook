package chapter4

import (
	"flag"
	"fmt"
	"time"
)

func FlagUsage() {
	var number int
	flag.IntVar(&number, "n", 10, "number")

	name := flag.String("name", "Go", "name of language")

	now := flag.Duration("time", time.Second, "time")
	flag.Parse()

	flag.PrintDefaults()
	fmt.Println(number, *name, *now)
}
