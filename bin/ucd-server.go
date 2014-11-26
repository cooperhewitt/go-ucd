package main

import(
	"fmt"
	_ "flag"
	"net/http"
	"org.cooperhewitt/ucd/names"
)

func string (w http.ResponseWriter, r *http.Request){

     string := r.FormValue("string")
     names := names.NamesForString(char)
     fmt.Fprintf(w, names)
}

func char (w http.ResponseWriter, r *http.Request){

     char := r.FormValue("char")
     name := names.Name(char)
     fmt.Fprintf(w, name)
}

func main(){
    
    http.HandleFunc("/char", char)
    http.ListenAndServe(":8080", nil)
}
