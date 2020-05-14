package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, `<h1>hello webserver</h1>`)
	})

	http.ListenAndServe(":8000", nil)
}

// $env GOOS=linux GOARCH=amd64 go build
