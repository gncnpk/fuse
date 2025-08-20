package handlers

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"
)

var kojiProUrl = "https://mgstpp-game.konamionline.com/"

func ToKojiPro(URL string, body io.Reader, length int64) (*http.Response, error) {
	c := http.Client{Timeout: time.Second * 10, Transport: &http.Transport{
		DisableCompression: true,
	}}

	newUrl, err := url.JoinPath(kojiProUrl, URL)
	if err != nil {
		return nil, fmt.Errorf("cannot make an url from %s: %w", URL, err)
	}

	req, err := http.NewRequest(http.MethodPost, newUrl, body)
	if err != nil {
		return nil, fmt.Errorf("cannot create request to kojipro: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Connection", "Keep-Alive")

	req.Header.Set("User-Agent", "")
	req.Header.Del("Transfer-Encoding")
	req.Header.Del("Accept-Encoding")

	req.ContentLength = length

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot make a request to %s: %w", newUrl, err)
	}

	slog.Debug("resp", "code", resp.StatusCode, "content-length", resp.ContentLength, "transfer-encoding", resp.TransferEncoding, "uncompressed", resp.Uncompressed)
	headers := ""
	for k, v := range resp.Header {
		headers += fmt.Sprintf("%s: %s ::: ", k, v)
	}
	slog.Debug("headers", "values", headers)

	return resp, nil
}
