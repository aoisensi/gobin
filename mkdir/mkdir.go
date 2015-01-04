package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	for _, path := range flag.Args() {
		err := os.Mkdir(path, 0777)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
