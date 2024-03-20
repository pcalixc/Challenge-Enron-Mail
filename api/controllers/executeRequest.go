package controllers

import (
	"challenge/api/config"
	"fmt"
	"net/http"
	"strings"
)

func executeRequest(query string) (*http.Response, error) {
	req, err := http.NewRequest("POST", config.GetEnvVar("ZS_BASE_URL")+"/api/emails/_search", strings.NewReader(query))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.SetBasicAuth(config.GetEnvVar("ZS_USER"), config.GetEnvVar("ZS_PASSWORD"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}

	return resp, nil
}
