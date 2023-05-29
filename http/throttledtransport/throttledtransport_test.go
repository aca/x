package throttledtransport_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/aca/x/http/throttledtransport"
)

func TestThrottledTransport(t *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "hello")
    }))
    defer ts.Close()

    req, _ := http.NewRequest(http.MethodGet, ts.URL, nil)
    client := http.Client{
        Transport: throttledtransport.NewThrottledTransport(time.Second * 5, 1, http.DefaultTransport),
    }

    t1 := time.Now()

    for i:=0; i< 4; i++{
        t.Logf("request %v start", i)
        client.Do(req)
        t.Logf("request %v end", i)
    }

    t2 := time.Now()
    t.Log(t2.Sub(t1))
}

