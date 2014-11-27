package names

import (
	"fmt"
	"org.cooperhewitt/ucd/data/unicodedata"
	"org.cooperhewitt/ucd/data/unihan"
	"strings"
	"unicode/utf8"
)

type UCDName struct {
	Char string
	Hex  string
	Name string
}

func (u UCDName) String() string {
	return u.Name
}

/*
   things that don't work because https://github.com/cooperhewitt/go-ucd/issues/1
   http://www.fileformat.info/info/unicode/char/4355/index.htm
*/

func Name(char string) (f UCDName) {

	hex := CharToHex(char, "unicodedata")
	name, _ := unicodedata.UCD[hex]

	if name == "" {
	   hex := CharToHex(char, "unihan")
	   name, _ = unihan.UCDHan[hex]
	}

	return UCDName{char, hex, name}
}

func NamesForString(s string) (n []UCDName) {

	chars := strings.Split(s, "")
	count := len(chars)

	results := make([]UCDName, count)

	for idx, char := range chars {
		name := Name(char)
		results[idx] = name
	}

	return results
}

func CharToHex(char string, table string) (hex string) {

	rune, _ := utf8.DecodeRuneInString(char)

	hex_fmt := "%04X"

	if table == "unihan" {
	   hex_fmt = "%05X"
	}	   
	
	hex = fmt.Sprintf(hex_fmt, rune)
	return hex
}
