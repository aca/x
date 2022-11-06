package httpsredirect

import (
	"log"
	"net/http"
)

func HTTPSRedirectHandler(w http.ResponseWriter, r *http.Request) {
	toURL := "https://"
	toURL += r.Host
	toURL += r.URL.RequestURI()
    log.Print(toURL)

	// w.Header().Set("Cache-Control", "public, max-age=2592000")
	w.Header().Set("Location", toURL)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(301)
	// http.Redirect(w, r, toURL, http.StatusMovedPermanently)
}
