package main

import (
	"io"
	"net/http"

	"github.com/aca/x/hsts"
)

func main() {
    f := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "helloworld")
	})

	http.Handle("/", hsts.HSTSMiddleware(f))

	http.ListenAndServe(":80", nil)
}
