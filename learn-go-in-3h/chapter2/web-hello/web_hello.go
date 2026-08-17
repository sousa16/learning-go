package main

import (
	"fmt"
	"net/http"
)

/*
	All Go programs start running from a function called `main` in a package
*/

func main() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		fmt.Fprintf(rw, "Hello, %s", name)
	})
	http.ListenAndServe(":8080", nil)
}
