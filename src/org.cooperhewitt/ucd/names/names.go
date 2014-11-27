package names

import (
       "fmt"
       "strings"
       "unicode/utf8"
       "org.cooperhewitt/ucd/data"
)

type UCDName struct {
     Char string
     Name string
}

func Name(char string) (f UCDName){

     hex := CharToHex(char)

     name, _ := ucd.UCD[hex]
     return UCDName{char, name}
}

func NamesForString(s string) (n []UCDName){

     count := len(s)

     results := make([]UCDName, count)

     chars := strings.Split(s, "")

     for idx, char := range chars {
          name := Name(char)	 
	  results[idx] = name
     }

     return results
}

func CharToHex(char string) (hex string){

     rune, _ := utf8.DecodeRuneInString(char)
     hex = fmt.Sprintf("%04X", rune)

     return hex
}
