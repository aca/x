package chromecookies

import (
	"context"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

// SetCookiesFromRemoteAllocator(http.DefaultClient, "ws://127.0.0.1:9222/")
func SetCookiesFromRemoteAllocator(httpc *http.Client, remoteURL string) {
	actx, cancel := chromedp.NewRemoteAllocator(context.Background(), "ws://127.0.0.1:9222/")
	defer cancel()
	actx, cancel = chromedp.NewContext(actx)
	defer cancel()

	if httpc.Jar == nil {
		jar, _ := cookiejar.New(nil)
		httpc.Jar = jar
	}

	chromedp.Run(actx,
		chromedp.Tasks{
			chromedp.ActionFunc(func(ctx context.Context) error {
				cookies, err := network.GetAllCookies().Do(ctx)
				if err != nil {
					log.Fatal(err)
				}

				cookieMap := make(map[string][]*http.Cookie)

				for _, cookie := range cookies {
					cookieMap[cookie.Domain] = append(cookieMap[cookie.Domain], &http.Cookie{
						Name:     cookie.Name,
						Value:    cookie.Value,
						Path:     cookie.Path,
						Domain:   cookie.Domain,
						Secure:   cookie.Secure,
						HttpOnly: cookie.HTTPOnly,
					})
				}

				for k, v := range cookieMap {
					{
						u := k
						u = "http://" + k
						urlObj, err := url.Parse(u)
						if err != nil {
							log.Fatal(err)
						}
						httpc.Jar.SetCookies(urlObj, v)
					}

					{
						u := k
						u = "https://" + k
						urlObj, err := url.Parse(u)
						if err != nil {
							log.Fatal(err)
						}
						httpc.Jar.SetCookies(urlObj, v)
					}

				}
				return nil
			},
			),
		},
	)
}
