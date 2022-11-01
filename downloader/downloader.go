package downloader

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type Downloader struct {
	Client  *http.Client
	Headers [][]string
}

var DefaultDownloader Downloader = Downloader{
	Client: http.DefaultClient,
	Headers: [][]string{
		{"user-agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36`},
	},
}

func (d *Downloader) Download(rawURL string, output string) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("failed to parse %s: %w", u, err)
	}

	if output == "" {
		output = filepath.Join(u.Host, u.Path)
	}

	if fileExists(output) {
		return nil
	}

	err = os.MkdirAll(filepath.Dir(output), 0o777)
	if err != nil {
		return fmt.Errorf("failed to download %s: %w", u, err)
	}

	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
    if err != nil {
        return err
    }

	for _, header := range d.Headers {
		if len(header) != 2 {
			return fmt.Errorf("invalid header k/v pair %+v", header)
		}
		req.Header.Set(header[0], header[1])
	}

	resp, err := d.Client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download %s: %w", u, err)
	}
	defer resp.Body.Close()

	f, err := os.Create(output + ".tmp")
	if err != nil {
		return fmt.Errorf("failed to download %s: %w", u, err)
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to download %s: %w", u, err)
	}

	err = os.Rename(output+".tmp", output)
	if err != nil {
		return fmt.Errorf("failed to download %s: %w", u, err)
	}

	return nil
}

func Download(url string, output string) error {
	return DefaultDownloader.Download(url, output)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
