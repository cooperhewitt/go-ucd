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

The following tools are included with this repository. _Note however that you will need to [compile them](https://golang.org/cmd/go/#hdr-Compile_and_run_Go_program) yourself._

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

`ucd` supports the Unicode Han Data character set, or at least endeavours to. There may still be bugs.

	$> ucd 䍕
	NET; WEB; NETWORK, NET FOR CATCHING RABBIT

### ucd-server

#### Usage

	$> ucd-server --help
	Usage of ./ucd-server:
	  -host="localhost": host
	  -port=8080: port

#### as JSON

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

#### As plain text

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

## Versions

This package exports data defined in the `UnicodeData.txt` and the `Unihan.zip`. Both are available from
http://unicode.org/Public/UCD/latest/ucd/. You can see the dates that each was
last compiled in to the source code for `ucd` at the top of each source file in
[the data directory](https://github.com/cooperhewitt/go-ucd/tree/master/src/org.cooperhewitt/ucd/data).

## Shout outs

Many thanks to friend and Go-friend [Richard Crowley](https://github.com/rcrowley) who is always kind and patient answering my Go-related questions. Go is lovely but Go is weird.

## See also

* http://unicode.org/Public/UCD/latest/ucd/
* http://www.washingtonpost.com/news/the-intersect/wp/2015/02/23/the-surprisingly-complex-reason-you-never-see-emoji-urls/
* https://modelviewculture.com/pieces/i-can-text-you-a-pile-of-poo-but-i-cant-write-my-name
