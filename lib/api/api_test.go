package api

import (
	"net/http"
	"testing"
)

func TestSendHTTPRequest(t *testing.T) {
	tests := []struct {
		name           string
		httpMethod     string
		apiEndpoint    string
		resourcePath   string
		requestHeaders map[string]string
		queries        []string
		requestBody    string
		expectedError  string
	}{
		{
			name:          "Invalid schema",
			httpMethod:    http.MethodGet,
			apiEndpoint:   "htt//example.com",
			resourcePath:  "/test",
			queries:       []string{"param1:value1", "param2:value2"},
			expectedError: "Get \"/test?param1=value1&param2=value2\": unsupported protocol scheme \"\"",
		},
		{
			name:          "Valid GET Request",
			httpMethod:    http.MethodGet,
			apiEndpoint:   "https://example.com",
			resourcePath:  "/test",
			queries:       []string{"param1:value1", "param2:value2"},
			expectedError: "",
		},
		{
			name:          "Invalid HTTP Method",
			httpMethod:    "INVALID",
			apiEndpoint:   "https://example.com",
			resourcePath:  "/test",
			expectedError: "invalid HTTP method: INVALID",
		},
		{
			name:          "Invalid Query Format",
			httpMethod:    http.MethodGet,
			apiEndpoint:   "https://example.com",
			resourcePath:  "/test",
			queries:       []string{"invalidquery"},
			expectedError: "invalid query format: invalidquery",
		},
		{
			name:           "Valid POST Request",
			httpMethod:     http.MethodPost,
			apiEndpoint:    "https://example.com",
			resourcePath:   "/test",
			requestHeaders: map[string]string{"Content-Type": "application/json"},
			requestBody:    `{"key": "value"}`,
			expectedError:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := SendHTTPRequest(tt.httpMethod, tt.apiEndpoint, tt.resourcePath, tt.requestHeaders, tt.queries, tt.requestBody)

			if tt.expectedError != "" {
				if err == nil || err.Error() != tt.expectedError {
					t.Errorf("Expected error: %s, got: %v", tt.expectedError, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if res == nil {
					t.Error("Expected non-nil response, got nil")
				}
			}
		})
	}
}
