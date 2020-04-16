package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// DispatchRequest returns response received from target resource
func DispatchRequest(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r) // Need to embed optional slash parsinfg in URL
	params, ok := r.URL.Query()["u"]
	if !ok || len(params) == 0 {
		fmt.Fprintf(w, "Target URL invalid or not specified.")
	}
	targetURL := params[0]

	fmt.Println("PING", targetURL) // DEV
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

	// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, `To apply this proxy, suffix the desired URL to this
	// 	URL. However, this may throw an error if the URL is invalid.`)
	// }).Methods("GET")

	r.HandleFunc("/", DispatchRequest).Methods("GET")

	fmt.Println("Listeing...")
	log.Fatal(http.ListenAndServe(":8001", r))
}
