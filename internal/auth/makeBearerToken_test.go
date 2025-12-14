package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		wantToken   string
		expectError bool
	}{
		{
			name: "valid bearer token",
			headers: http.Header{
				"Authorization": []string{"Bearer abc123"},
			},
			wantToken:   "abc123",
			expectError: false,
		},
		{
			name: "missing Authorization header",
			headers: http.Header{},
			expectError: true,
		},
		{
			name: "incorrect prefix",
			headers: http.Header{
				"Authorization": []string{"Token abc123"},
			},
			expectError: true,
		},
		{
			name: "bearer prefix but empty token",
			headers: http.Header{
				"Authorization": []string{"Bearer "},
			},
			expectError: true,
		},
		{
			name: "bearer prefix without space",
			headers: http.Header{
				"Authorization": []string{"Bearer"},
			},
			expectError: true,
		},
		{
			name: "extra whitespace after bearer",
			headers: http.Header{
				"Authorization": []string{"Bearer   abc123"},
			},
			wantToken:   "  abc123",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GetBearerToken(tt.headers)

			if tt.expectError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if token != tt.wantToken {
				t.Errorf("expected token %q, got %q", tt.wantToken, token)
			}
		})
	}
}
