package main

import(
	"fmt"
	_ "flag"
	"net/http"
	"org.cooperhewitt/ucd/names"
)

func ucd (w http.ResponseWriter, r *http.Request){

     str := r.FormValue("string")
     name := names.Name(str)
     fmt.Fprintf(w, name)
}

func main(){
    
    http.HandleFunc("/", ucd)
    http.ListenAndServe(":8080", nil)
}
