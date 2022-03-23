package main

import (
	"fmt"
	"net/http"
)

var PORT = ":8181"

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello Golang"
	fmt.Fprintf(w, msg)
}
