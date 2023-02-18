package util

import (
	"context"
	"net/http"
	"time"
)

func HttpGetWithTimeout(url string, timeout time.Duration) (*http.Response, error) {
	cntx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	client := &http.Client{}
	request, _ := http.NewRequestWithContext(
		cntx,
		http.MethodGet,
		url,
		nil)
	return client.Do(request)
}
