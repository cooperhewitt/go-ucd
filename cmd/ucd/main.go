package main

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-ucd/v13"
	"strings"
)

func main() {

	flag.Parse()

	args := flag.Args()
	str := strings.Join(args, " ")

	chars := strings.Split(str, "")

	for _, char := range chars {
		n := ucd.Name(char)
		fmt.Println(n)
	}
}
