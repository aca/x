package throttledtransport

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type ThrottledTransport struct {
	roundTripper http.RoundTripper
	ratelimiter      *rate.Limiter
}

func (t *ThrottledTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	err := t.ratelimiter.Wait(r.Context())
	if err != nil {
		return nil, err
	}
	return t.roundTripper.RoundTrip(r)
}

func NewThrottledTransport(interval time.Duration, burst int, rt http.RoundTripper) http.RoundTripper {
	return &ThrottledTransport{
		roundTripper: rt,
		ratelimiter:      rate.NewLimiter(rate.Every(interval), burst),
	}
}
