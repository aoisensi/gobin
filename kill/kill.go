package main

import (
	"flag"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		pid, err := strconv.Atoi(arg)
		if err != nil {
			os.Exit(-1)
		}
		p, err := os.FindProcess(pid)
		if err != nil {
			os.Exit(-1)
		}
		p.Kill()
	}
}
