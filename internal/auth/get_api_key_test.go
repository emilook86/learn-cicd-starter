package auth

import (
	"net/http"
	"testing"
)

func TestAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey secret-token-123")

	apiKey, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("expected %v, got %v", nil, err)
	}

	if apiKey != "secret-token-123" {
		t.Fatalf("expected %v, got %v", "secret-token-123", apiKey)
	}
}

func TestAPIKeyNoAuth(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)

	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected %v, got %v", ErrNoAuthHeaderIncluded, err)
	}
}
