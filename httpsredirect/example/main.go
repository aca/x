package main

import (
	"net/http"

	"github.com/aca/x/httpsredirect"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        httpsredirect.HTTPSRedirectHandler(w, r)
    })

    http.ListenAndServe(":80", nil)
}
