package names

import (
       "strings"
       "fmt"
       "unicode/utf8"
       "org.cooperhewitt/ucd/data"
)

func Name(k string) (v string){
     h := ToHex(k)
     return ucd.UCD[h]
}

func ToHex(s string) (h string){

     // seriously why is it so hard to do the
     // equivalent of python's this:     	    
     // id = ord(char)
     // hex = "%04X" % id

     s = fmt.Sprintf("%+q", s)

     s = strings.Replace(s, "\"", "", 2)
     // fmt.Printf("s becomes '%v'\n", s)
     // fmt.Printf("s is '%c'\n", s)

     c := utf8.RuneCountInString(s)
     
     s = s[c-4:c]
     s = strings.ToUpper(s)    
     return s
}
