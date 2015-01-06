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
)

func init() {
	flag.BoolVar(&flagAlmostAll, "A", false, "Almost All")
	flag.BoolVar(&flagLong, "l", false, "Show All Info (same -P)")
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
}

func toMessage(file os.FileInfo) []string {
	var result []string
	if flagLong {
		result = []string{
			parsePerm(file),
			parseTime(file),
			file.Name(),
		}
	} else {
		result = []string{file.Name()}
	}
	return result
}

func parsePerm(file os.FileInfo) string {
	r := file.Mode().Perm().String()
	if file.Mode().IsDir() {
		r = "d" + r[1:]
	}
	return r
}

func parseTime(file os.FileInfo) string {
	r := file.ModTime()
	return r.String()
}
