package main

import(
	"fmt"
	"flag"
	"org.cooperhewitt/ucd"
)

func main(){

     flag.Parse()
     char := flag.Arg(0)

     name := ucd.Name(char)
     fmt.Println(name)
}
