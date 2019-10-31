package main

import (
	"fmt"
	"net/http"

	"rsc.io/quote"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, quote.Hello())
	})
	http.ListenAndServe(":8080", nil)
}
