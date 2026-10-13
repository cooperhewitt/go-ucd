package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"maps"
	"os"

	"github.com/aaronland/go-ucd/v17/unicodedata"
	"github.com/aaronland/go-ucd/v17/unihan"
	"github.com/tidwall/pretty"
)

func main() {

	var include_unihan = flag.Bool("unihan", false, "Include Unihan data.")
	var make_pretty = flag.Bool("pretty", false, "Prettify JSON output.")

	flag.Parse()

	var dump map[string]string
	dump = unicodedata.UCD

	if *include_unihan {

		// sorting?

		maps.Copy(dump, unihan.UCDHan)
	}

	body, err := json.Marshal(dump)

	if err != nil {
		log.Fatal(err)
	}

	if *make_pretty {
		body = pretty.Pretty(body)
	}

	fmt.Printf("%s\n", body)
	os.Exit(0)
}
