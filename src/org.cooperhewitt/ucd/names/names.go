package names

import (
       "fmt"
       "unicode/utf8"
       "org.cooperhewitt/ucd/data"
)

func Name(char string) (name string){

     hex := CharToHex(char)

     name, _ = ucd.UCD[hex]
     return name
}

/*
func NamesForString(s string) (n array){

     chars := strings.Split(str, "")

     for _, char := range chars {
          name := names.Name(char)
	  
     }
}
*/

func CharToHex(char string) (hex string){

     rune, _ := utf8.DecodeRuneInString(char)
     hex = fmt.Sprintf("%04X", rune)

     return hex
}
