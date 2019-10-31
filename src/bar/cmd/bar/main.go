package main

import (
	"net/http"

	"github.com/drillbits/golangci-lint-test/src/bar"
)

func main() {
	http.HandleFunc("/", bar.Handler)
	http.ListenAndServe(":8888", nil)
}
