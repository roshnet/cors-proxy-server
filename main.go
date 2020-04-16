package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Mux!")
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")
	fmt.Println("Listening on localhost:5000")
	http.ListenAndServe(":5000", r)
}
