package apphttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

var client = &http.Client{}

func send(req *http.Request) (*http.Response, http.Header, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to send request to %s\nException: %w", req.RequestURI, err)
	}
	return resp, resp.Header, nil
}

func Post(requestUrl string, body interface{}, header http.Header, ctx context.Context) (*http.Response, http.Header, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal body: %w", err)
	}

	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create new request: %w", err)
	}

	if header != nil {
		req.Header = header
	}

	req.Header.Add("Content-Type", "application/json")

	if ctx != nil {
		req.WithContext(ctx)
	}

	return send(req)
}
