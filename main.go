package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// DispatchRequest returns response received from target resource
func DispatchRequest(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r) // Need to embed optional slash parsinfg in URL
	params := r.URL.Query()
	targetURL := params.Get("u")

	if targetURL == "" {
		fmt.Fprintf(w, "Target URL not specified.")
		log.Fatal("Exiting request. URL not specified.")
		return
	}

	// TBD: Validate `targetURL` structure in preflight request (1).
	// TBD: Allow more params for customized functionality (2).

	response, err := http.Get(targetURL)
	if err != nil {
		panic(err) // Maybe use log.Fatal() instead?
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	fmt.Fprintln(w, string(body))
}

func main() {
	r := mux.NewRouter()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
Usage: https://proxify-cors.herokuapp.com/proxy?u=&lt;desired&gt; resource
absolute path. However, this may throw an error if the URL is invalid.`)
	}).Methods("GET")

	r.HandleFunc("/proxy", DispatchRequest).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, r))
}
