package util

import (
	"context"
	"net/http"
)

func HttpGet(cntx context.Context, url string) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequestWithContext(
		cntx,
		http.MethodGet,
		url,
		nil)
	return client.Do(request)
}
