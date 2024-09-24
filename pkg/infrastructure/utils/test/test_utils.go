package test

import (
	"context"
	"io"
	"net/http"
)

func ExecuteRequest(url string, body io.Reader, methodHTTP string) (*http.Response, error) {
	cli := &http.Client{}
	req, _ := http.NewRequestWithContext(context.Background(), methodHTTP, url, body)
	req.Header.Set("Content-Type", "application/json")
	return cli.Do(req)
}
