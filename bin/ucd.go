package main

import(
	"strings"
	"fmt"
	"flag"
	"org.cooperhewitt/ucd/names"
)

func main(){

     flag.Parse()

     args:= flag.Args()
     str := strings.Join(args, " ")

     chars := strings.Split(str, "")

     for _, char := range chars {
          name := names.Name(char)
     	  fmt.Println(name)
     }	  
}
