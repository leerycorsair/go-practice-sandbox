package main

import (
	"errors"
	"testing"

	"github.com/beevik/ntp"
)

type MockNTPClient struct {
	Response *ntp.Response
	Err      error
}

func (m *MockNTPClient) Query(server string) (*ntp.Response, error) {
	return m.Response, m.Err
}

func TestProcessTime(t *testing.T) {
	tests := []struct {
		name        string
		client      NTPClient
		expectError string
	}{
		{
			name:        "query error",
			client:      &MockNTPClient{Err: errors.New("query error")},
			expectError: "query error",
		},
		{
			name:        "response validation error",
			client:      &MockNTPClient{Response: &ntp.Response{}, Err: errors.New("kiss of death recieved")},
			expectError: "kiss of death recieved",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ProcessTime(tt.client, "0.beevik-ntp.pool.ntp.org")
			if (err != nil) && (err.Error() != tt.expectError) {
				t.Errorf("expected error %v, got %v", tt.expectError, err.Error())
			}
			if err == nil && tt.expectError != "" {
				t.Errorf("expected error %v, got nil", tt.expectError)
			}
		})
	}
}
