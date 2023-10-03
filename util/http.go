package util

import (
	"context"
	"io"
	"net"
	"net/http"
	"time"
)

// read more at: https://pkg.go.dev/net/http#DefaultTransport
var httpClient http.Client = http.Client{
	Transport: &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          10,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		IdleConnTimeout:       10 * time.Second,
		ExpectContinueTimeout: 0,
	},
}

func HttpGet(ctx context.Context, url string) (*http.Response, error) {
	request, _ := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil)
	return http.DefaultClient.Do(request)
}

func HttpGetJson[T any](ctx context.Context, url string) (*T, error) {
	var result T
	request, _ := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil)

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
