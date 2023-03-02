package hsts

import (
	"net/http"
	"strconv"
	"time"
)

var (
	maxAge              = time.Hour * 24 * 126
	maxAgeSeconds       = int(maxAge.Seconds())
	maxAgeSecondsString = strconv.Itoa(maxAgeSeconds)
	headerValue         = "max-age=" + maxAgeSecondsString + "; includeSubDomains"
)

func HSTSMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age="+strconv.Itoa(maxAgeSeconds)+"; includeSubDomains")
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
