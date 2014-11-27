package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"org.cooperhewitt/ucd/names"
)

type UCDResponse struct {
     Chars []names.UCDName
}

func string(w http.ResponseWriter, r *http.Request) {

	txt := r.FormValue("text")
	chars := names.NamesForString(txt)

	rsp := UCDResponse{chars}
	send(w, r, rsp)
}

func char(w http.ResponseWriter, r *http.Request) {

	txt := r.FormValue("text")
	char := names.Name(txt)

	chars := make([]names.UCDName, 1)
	chars[0] = char

	rsp := UCDResponse{chars}
	send(w, r, rsp)
}

func send(w http.ResponseWriter, r *http.Request, rsp UCDResponse) {

	send_json(w, rsp)
}

/*
func send_text(w http.ResponseWriter, rsp UCDResponse) {

     w.Header().Set("Content-Type", "application/json")

     for _, char := range rsp.Chars {
     	 fmt.Fprintf(w ,char)
     }     
}
*/

func send_json(w http.ResponseWriter, rsp UCDResponse) {

	js, err := json.Marshal(rsp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
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
