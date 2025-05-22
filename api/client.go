package api

import (
	"net/http"
	"time"
)

// Declare client struct
type Client struct {
	httpClient http.Client
}

// Create a new client
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
