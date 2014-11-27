package main

import (
	"flag"
	"fmt"
	"org.cooperhewitt/ucd/names"
	"strings"
)

func main() {

	flag.Parse()

	args := flag.Args()
	str := strings.Join(args, " ")

	chars := strings.Split(str, "")

	for _, char := range chars {
		ucd := names.Name(char)
		fmt.Println(ucd.Char)
	}
}
