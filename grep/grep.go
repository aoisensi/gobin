package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		os.Exit(-1)
	}
	word := args[0]
	files := args[1:]
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
			text := scanner.Text()
			if strings.Index(text, word) != -1 {
				fmt.Println(scanner.Text())
			}
		}
	}
}
