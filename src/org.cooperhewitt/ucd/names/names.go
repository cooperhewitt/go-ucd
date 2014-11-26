package names

import (
       "fmt"
       "strings"
       "unicode/utf8"
       "org.cooperhewitt/ucd/data"
)

func Name(char string) (name string){

     hex := CharToHex(char)

     name, _ = ucd.UCD[hex]
     return name
}

func NamesForString(s string) (n [][]string){

     chars := strings.Split(s, "")
     count := len(chars)

     var results [4][2]string

     i := 0

     for _, char := range chars {
          name := Name(char)
	  
	  var lookup [2]string
	  lookup[0] = char
	  lookup[1] = name

	  results[i] = lookup
	  i += 1
     }

     return results
}

func CharToHex(char string) (hex string){

     rune, _ := utf8.DecodeRuneInString(char)
     hex = fmt.Sprintf("%04X", rune)

     return hex
}
