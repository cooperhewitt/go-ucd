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

     count := len(s)

     results := make([][]string, count)

     for i := range results {
     	 results[i] = make([]string, 2)
     }

     chars := strings.Split(s, "")

     for idx, char := range chars {
          name := Name(char)
	  
	  results[idx][0] = char
	  results[idx][1] = name
     }

     return results
}

func CharToHex(char string) (hex string){

     rune, _ := utf8.DecodeRuneInString(char)
     hex = fmt.Sprintf("%04X", rune)

     return hex
}
