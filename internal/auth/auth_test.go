package auth

import (
	"context"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://test.com", nil)
	if err != nil {
		t.Fatalf("unexpected error creating request: %v", err)
	}

	testCases := []struct {
		description   string
		authHeader    string
		expectedKey   string
		expectedError bool
	}{
		{
			description:   "CorrectAPIKey",
			authHeader:    "ApiKey expected-key",
			expectedKey:   "good-key",
			expectedError: false,
		},
		{
			description:   "BadAuthorizationScheme",
			authHeader:    "Bearer bad-key",
			expectedError: true,
		},
		{
			description:   "EmptyAuthorizationHeader",
			authHeader:    "",
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			req.Header.Set("Authorization", tc.authHeader)
			apiKey, err := GetAPIKey(req.Header)
			if tc.expectedError {
				if err == nil {
					t.Error("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error getting API Key: %v", err)
				}
				if apiKey != "expected-key" {
					t.Errorf("expected API key to be %q, but got %q", "ApiKey", apiKey)
				}
			}
		})
	}

}
