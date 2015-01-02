package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	flagA bool
)

func init() {
	flag.BoolVar(&flagA, "A", false, "all")
}

func main() {
	flag.Parse()
	paths := flag.Args()
	if len(paths) == 0 {
		paths = []string{"."}
	}
	path := paths[0]
	files, err := ioutil.ReadDir(path)
	if err != nil {
		os.Exit(-1)
	}
	for _, file := range files {
		name := file.Name()
		if flagA || name[0] != '.' {
			fmt.Println(file.Name())
		}
	}
}
