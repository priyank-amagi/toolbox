package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SendHTTPRequest sends a generic HTTP request.
func SendHTTPRequest(httpMethod, apiEndpoint, resourcePath string, requestHeaders map[string]string, queries []string, requestBody string) (*http.Response, error) {
	// Validate HTTP method
	method := http.MethodGet // Default to GET
	switch httpMethod {
	case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete:
		method = httpMethod
	default:
		return nil, fmt.Errorf("invalid HTTP method: %s", httpMethod)
	}

	// Validate and build URL
	endpointURL, err := url.Parse(apiEndpoint)
	if err != nil {
		return nil, err
	}
	endpointURL.Path = resourcePath

	// Add queries to URL
	queryParams := endpointURL.Query()
	for _, query := range queries {
		parts := strings.SplitN(query, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid query format: %s", query)
		}
		queryParams.Add(parts[0], parts[1])
	}
	endpointURL.RawQuery = queryParams.Encode()

	// Create HTTP request
	request, err := http.NewRequest(method, endpointURL.String(), strings.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	// Add headers to request
	for key, value := range requestHeaders {
		request.Header.Add(key, value)
	}

	client := http.DefaultClient
	return client.Do(request)
}
