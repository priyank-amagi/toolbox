package todos

import (
	"net/http"
	"testing"

	"github.com/priyank-amagi/toolbox/lib/http_helper"
	http_mocks "github.com/priyank-amagi/toolbox/lib/http_helper/mocks"
	"gotest.tools/assert"
)

func TestNewTodoSvc(t *testing.T) {
	httpSvc := http_helper.NewHttpSvc()
	testHost := "test_json_host_name"
	assert.Equal(t, *NewTodoSvc(testHost, httpSvc), TodoSvc{hostname: testHost, httpSvc: httpSvc})
}

// TODO: this tests are incomplete
func TestGetTodos(t *testing.T) {
	httpSvc := http_mocks.NewIHttpSvc(t)
	tests := []struct {
		name          string
		host          string
		todos         []Todo
		request       *http.Request
		response      *http.Response
		expectedError string
	}{
		{
			name:          "Invalid URL",
			host:          "://invalid-url",
			expectedError: "parse \"://invalid-url\": missing protocol scheme",
		},
	}
	for _, tt := range tests {
		todoSvc := NewTodoSvc(tt.host, httpSvc)
		t.Run(tt.name, func(t *testing.T) {
			todos, err := todoSvc.GetTodos()
			if tt.expectedError != "" {
				if err == nil || err.Error() != tt.expectedError {
					t.Errorf("Expected error: %s, got: %v", tt.expectedError, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				assert.Equal(t, tt.todos, todos)
			}
		})
	}
}
