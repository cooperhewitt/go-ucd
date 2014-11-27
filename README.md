# go-ucd

Go libraries and utilities for working with Unicode character data.

## Example

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

## bin

The following tools are included with this repository. _Note however that you will need to `go build` them yourself._

### ucd

	$> ucd A
	LATIN CAPITAL LETTER A

	$> ucd THIS → WAY
	LATIN CAPITAL LETTER T
	LATIN CAPITAL LETTER H
	LATIN CAPITAL LETTER I
	LATIN CAPITAL LETTER S
	SPACE
	RIGHTWARDS ARROW
	SPACE
	LATIN CAPITAL LETTER W
	LATIN CAPITAL LETTER A
	LATIN CAPITAL LETTER Y

### ucd-server

	$> ucd-server --help
	Usage of ./ucd-server:
	  -host="localhost": host
	  -port=8080: port

	$> curl -X GET -s 'http://localhost:8080/?text=♕%20HAT' | python -mjson.tool
	{
	    "Chars": [
	        {
	            "Char": "\u2655",
	            "Hex": "2655",
	            "Name": "WHITE CHESS QUEEN"
	        },
	        {
	            "Char": " ",
	            "Hex": "0020",
	            "Name": "SPACE"
	        },
	        {
	            "Char": "H",
	            "Hex": "0048",
	            "Name": "LATIN CAPITAL LETTER H"
	        },
	        {
	            "Char": "A",
	            "Hex": "0041",
	            "Name": "LATIN CAPITAL LETTER A"
	        },
	        {
	            "Char": "T",
	            "Hex": "0054",
	            "Name": "LATIN CAPITAL LETTER T"
	        }
	    ]
	}

	$> curl -H 'Accept: text/plain' -s 'http://localhost:8080/?text=♕%20HAT%20WITH%20😸'
	WHITE CHESS QUEEN
	SPACE
	LATIN CAPITAL LETTER H
	LATIN CAPITAL LETTER A
	LATIN CAPITAL LETTER T
	SPACE
	LATIN CAPITAL LETTER W
	LATIN CAPITAL LETTER I
	LATIN CAPITAL LETTER T
	LATIN CAPITAL LETTER H
	SPACE
	GRINNING CAT FACE WITH SMILING EYES

## Known knowns

* Things like the `Unicode Han Data` character set are not supported because they are not included in the Unicode Consortium's default [UnicodeData.txt](http://unicode.org/Public/UCD/latest/ucd/UnicodeData.txt) file.

## Shout outs

Many thanks to friend and Go-friend [Richard Crowley](https://github.com/rcrowley) who is always kind and patient answering my Go-related questions. Go is lovely but Go is weird.

## See also

* https://github.com/cooperhewitt/plumbing-ucd-server
* http://unicode.org/Public/UCD/latest/ucd/
