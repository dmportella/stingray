package main

import (
	"fmt"
	"time"
)

// NewClient Creates a new instance of the stingray api client.
func NewClient(config config) *Client {
	return &Client{
		config:    config,
		userAgent: generateUserAgent(config),
	}
}

func generateUserAgent(config Config) string {
	if config && config.userAgent != "" {
		return fmt.Sprintf(userAgentTemplate, config.userAgent)
	}

	return fmt.Sprintf(userAgentTemplate, "application name not set")
}

func (client *Client) createRequest() {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))

	req.Header.Set("accept", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", client.UserAgent)

	httpClient := &http.Client{Timeout: (120 * time.Second)}
}
