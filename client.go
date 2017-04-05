package stingray

import (
	"net/http"
	"time"
)

const (
	userAgentTemplate = "Stingray API (https://github.com/dmportella/stingray, 0.0.0, '%s')"
)

// Client The Api client for stingray.
type Client struct {
	config        Config
	userAgent     string
	httpTimeout   time.Duration
	httpTransport *http.Transport
}

// Config The confid structure for the client api for stingray.
type Config struct {
	UserAgent          string        `json:"userAgent"`
	HTTPTimeout        time.Duration `json:"duration"`
	HTTPEndpoint       string        `json:"endpoint"`
	InsecureSkipVerify bool          `json:"insecureSkipVerify"`
	UserName           string        `json:"username"`
	Password           string        `json:"password"`
}
