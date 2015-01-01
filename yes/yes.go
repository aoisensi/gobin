package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := strings.Join(os.Args[1:], " ")
	if msg == "" {
		msg = "y"
	}
	for {
		fmt.Println(msg)
	}
}
