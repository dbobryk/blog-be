package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	

	r := mux.NewRouter()
	r.HandleFunc("/health", health)

	http.ListenAndServe(":8080", r)
}

func health(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "All is okay.\n")

}
