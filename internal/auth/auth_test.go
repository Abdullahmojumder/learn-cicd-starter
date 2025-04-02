package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		{
			name:          "Valid API Key",
			headers:       http.Header{"Authorization": []string{"ApiKey my-secret-key"}},
			expectedKey:   "my-secret-key",
			expectedError: nil,
		},
		{
			name:          "Missing Authorization Header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name:          "Malformed Authorization Header",
			headers:       http.Header{"Authorization": []string{"Bearer my-secret-key"}},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name:          "Empty API Key",
			headers:       http.Header{"Authorization": []string{"ApiKey "}},
			expectedKey:   "",
			expectedError: nil, // No explicit error, but might want to add stricter validation
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.headers)

			if apiKey != tt.expectedKey {
				t.Errorf("expected key %q, got %q", tt.expectedKey, apiKey)
			}

			if (err != nil && tt.expectedError == nil) || (err == nil && tt.expectedError != nil) {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}

			if err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error %q, got %q", tt.expectedError, err)
			}
		})
	}
}
