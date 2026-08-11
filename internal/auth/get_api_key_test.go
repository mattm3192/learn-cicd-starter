package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		headers    http.Header
		wantKey    string
		wantErrMsg string
	}{
		{
			name:       "No Auth Header Included",
			headers:    http.Header{},
			wantKey:    "",
			wantErrMsg: "no authorization header included",
		},
		{
			name:       "Malformed Auth Header",
			headers:    http.Header{"Authorization": []string{"Bearer secret123"}},
			wantKey:    "",
			wantErrMsg: "malformed authorization header",
		},
		{
			name:       "valid test",
			headers:    http.Header{"Authorization": []string{"ApiKey secret123"}},
			wantKey:    "secret123",
			wantErrMsg: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api, err := GetAPIKey(tt.headers)
			if err != nil && !strings.Contains(err.Error(), tt.wantErrMsg) {
				t.Fatalf("expected: %v, got: %v", tt.wantErrMsg, err.Error())
			}
			if api != tt.wantKey {
				t.Fatalf("expected: %v, got: %v", tt.wantKey, api)
			}

		})
	}
}
