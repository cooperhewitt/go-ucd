package main

import (
	"encoding/json"
	_ "flag"
	"fmt"
	"net/http"
	"org.cooperhewitt/ucd/names"
)

func string(w http.ResponseWriter, r *http.Request) {

	string := r.FormValue("string")
	rsp := names.NamesForString(string)

	js, err := json.Marshal(rsp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func char(w http.ResponseWriter, r *http.Request) {

	char := r.FormValue("char")
	name := names.Name(char)
	fmt.Fprintf(w, name)
}

func main() {

	http.HandleFunc("/char", char)
	http.HandleFunc("/string", string)

	http.ListenAndServe(":8080", nil)
}
