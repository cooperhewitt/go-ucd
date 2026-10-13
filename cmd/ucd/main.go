package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/aaronland/go-ucd/v17"
)

func main() {

	flag.Parse()

	args := flag.Args()
	str := strings.Join(args, " ")

	chars := strings.SplitSeq(str, "")

	for char := range chars {
		n := ucd.Name(char)
		fmt.Println(n)
	}
}
