package util

import (
	"context"
	"io"
	"net/http"
)

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

	response, err := http.DefaultClient.Do(request)
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
