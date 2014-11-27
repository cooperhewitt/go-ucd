package names

import (
	"fmt"
	"org.cooperhewitt/ucd/data"
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

	hex := CharToHex(char)

	name, _ := ucd.UCD[hex]
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

func CharToHex(char string) (hex string) {

	rune, _ := utf8.DecodeRuneInString(char)
	hex = fmt.Sprintf("%04X", rune)

	return hex
}
