package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	flagAlmostAll bool
	flagLong      bool
	flagPerm      bool
)

func init() {
	flag.BoolVar(&flagAlmostAll, "A", false, "Almost All")
	flag.BoolVar(&flagLong, "l", false, "Show All Info (same -P)")
	flag.BoolVar(&flagPerm, "P", false, "Show File Permission")
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
		if flagAlmostAll || name[0] != '.' {
			mess := toMessage(file)
			fmt.Println(strings.Join(mess, " "))
		}
	}
}

func flagCheck() {
	if flagLong {
		flagPerm = true
	}
}

func toMessage(file os.FileInfo) []string {
	result := []string{file.Name()}
	if flagLong {
		result = append(parsePerm(file), result...)
	}
	return result
}

func parsePerm(file os.FileInfo) []string {
	r := file.Mode().Perm().String()
	if file.Mode().IsDir() {
		r = "d" + r[1:]
	}
	return []string{r}
}
