package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
	"strings"
)

type hasher interface {
	New() hash.Hash
}

var (
	cryptos map[string]func() hash.Hash
)

func init() {
	cryptos = map[string]func() hash.Hash{
		"md5":    func() hash.Hash { return md5.New() },
		"sha1":   func() hash.Hash { return sha1.New() },
		"sha256": func() hash.Hash { return sha256.New() },
		"sha224": func() hash.Hash { return sha256.New224() },
		"sha512": func() hash.Hash { return sha512.New() },
		"sha384": func() hash.Hash { return sha512.New384() },
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		os.Exit(-1)
	}
	typename := args[0]
	files := args[1:]
	lowed := strings.ToLower(typename)
	if lowed == "list" {
		for cn := range cryptos {
			fmt.Println(cn)
		}
		return
	}
	hasher := cryptos[lowed]
	if hasher == nil {
		fmt.Printf("%s is not supported crypto\n", args[0])
		os.Exit(-1)
	}
	for _, name := range files {
		file, err := os.Open(name)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer file.Close()
		r := checkCrypto(file, hasher())
		fmt.Printf("%s %s\n", r, name)
	}
}

func checkCrypto(file *os.File, h hash.Hash) string {
	io.Copy(h, file)
	return hex.EncodeToString(h.Sum(nil))
}
