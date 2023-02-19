package util

import (
	"context"
	"net/http"
)

func HttpGet(ctx context.Context, url string) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil)
	return client.Do(request)
}
