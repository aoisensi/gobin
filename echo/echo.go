package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	flagN bool
)

func init() {
	flag.BoolVar(&flagN, "n", false, "do not output the trailing newline")
}

func main() {
	flag.Parse()
	strs := flag.Args()
	str := strings.Join(strs, " ")

	fmt.Print(str)
	if !flagN {
		fmt.Println()
	}
}
