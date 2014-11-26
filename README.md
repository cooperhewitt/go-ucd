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

	$> ucd-server

	$> curl 'http://localhost:8080?string=ᚓ'
	OGHAM LETTER EADHADH        

## Shout outs

Many thanks to friend and Go-friend [Richard Crowley](https://github.com/rcrowley) who is also kind and patient answering my Go-related questions. Go is lovely but Go is weird.

## See also

* https://github.com/cooperhewitt/plumbing-ucd-server
* http://unicode.org/Public/UCD/latest/ucd/
