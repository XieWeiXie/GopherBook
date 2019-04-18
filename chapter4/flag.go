package chapter4

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
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

type Numbers struct {
	Num []int
}

func (n *Numbers) Set(value string) error {
	sList := strings.Split(value, "|")
	var num []int
	for _, i := range sList {
		in, _ := strconv.Atoi(i)
		num = append(num, in)
	}
	n.Num = num
	return nil
}

func (n *Numbers) String() string {
	return fmt.Sprintf("%#v", n.Num)
}

func FlagSpecial() {
	var n Numbers
	flag.Var(&n, "n", "number to parse")
	flag.Parse()
	fmt.Println(n.Num)
}
