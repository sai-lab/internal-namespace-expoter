package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "internal-namespace-expoter lives")
	})

	log.Fatal(http.ListenAndServe(":1192", nil))
}
