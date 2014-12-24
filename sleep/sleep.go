package main

import (
	"flag"
	"strconv"
	"time"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return
	}
	t, _ := strconv.Atoi(args[0])
	time.Sleep(time.Duration(t) * time.Second)
}
