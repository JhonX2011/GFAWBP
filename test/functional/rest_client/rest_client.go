package restclient

import (
	"context"
	"io"
	"net/http"
)

type IClient interface {
	DoGet(context.Context, string) (*http.Response, error)
	DoPost(context.Context, string, io.Reader, map[string][]string) (*http.Response, error)
}

type client struct{}

func New() IClient {
	return &client{}
}

func (client *client) DoGet(ctx context.Context, url string) (*http.Response, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (client *client) DoPost(ctx context.Context, url string, body io.Reader, headers map[string][]string) (*http.Response, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	request.Header = headers
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
