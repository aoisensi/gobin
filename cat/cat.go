package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()
	files := flag.Args()
	if len(files) == 0 {
		files = []string{"-"}
	}
	for _, file := range files {
		var r io.Reader
		if file == "-" {
			r = os.Stdin
		} else {
			var err error
			r, err = os.Open(file)
			if err != nil {
				panic(err)
			}
		}

		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}
