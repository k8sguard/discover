package main

import (
	"fmt"
	"log"
	"net/http"
)

var version = "0.0.1"

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "version: %s", version)
}

func main() {
	http.HandleFunc("/version", versionHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
