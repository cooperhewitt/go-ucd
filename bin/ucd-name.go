package main

import(
	"fmt"
	"flag"
	"org.cooperhewitt/ucd/names"
)

func main(){

     flag.Parse()
     char := flag.Arg(0)

     name := names.Name(char)
     fmt.Println(name)
}
