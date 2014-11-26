package names

import (
       "fmt"
       "unicode/utf8"
       "org.cooperhewitt/ucd/data"
)

func Name(k string) (v string){
     h := ToHex(k)

     v, _ = ucd.UCD[h]
     return v
}

/*
func NamesForString(s string) (n array){

     chars := strings.Split(str, "")

     for _, char := range chars {
          name := names.Name(char)
	  
     }
}
*/

func ToHex(s string) (h string){

     rune, _ := utf8.DecodeRuneInString(s)
     hex := fmt.Sprintf("%04X", rune)

     return hex
}
