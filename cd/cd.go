package main

import (
	"flag"
	"os"
)

func main() {
	flag.Parse()
	path := flag.Arg(0)
	if path == "" {
		path = "."
	}
	os.Chdir(path)
}
