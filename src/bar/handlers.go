package bar

import (
	"fmt"
	"net/http"

	"rsc.io/quote"
)

func Handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, quote.Hello())
}
