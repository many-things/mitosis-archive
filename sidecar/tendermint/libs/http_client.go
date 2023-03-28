package libs

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	// TODO: make this configurable
	Client = http.DefaultClient
}

func JSONPost(url string, body interface{}) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	return Client.Do(request)
}

func JSONGet(url string) (*http.Response, error) {
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return Client.Do(request)
}
