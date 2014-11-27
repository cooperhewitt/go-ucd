package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"org.cooperhewitt/ucd/names"
)

func string(w http.ResponseWriter, r *http.Request) {

	txt := r.FormValue("text")
	rsp := names.NamesForString(txt)

	js, err := json.Marshal(rsp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func char(w http.ResponseWriter, r *http.Request) {

	txt := r.FormValue("text")
	ucd := names.Name(txt)
	fmt.Fprintf(w, ucd.Name)
}

func main() {

	host := flag.String("host", "localhost", "host")
	port := flag.Int("port", 8080, "port")

	flag.Parse()

	endpoint := fmt.Sprintf("%s:%d", *host, *port)

	fmt.Printf("listening on %s\n", endpoint)

	http.HandleFunc("/", char)
	http.HandleFunc("/string", string)

	http.ListenAndServe(endpoint, nil)
}
