package main

import (
	"time"
)

const (
	userAgentTemplate = "Stingray API (https://github.com/dmportella/stingray, 0.0.0, '%s')"
)

// Client The Api client for stingray.
type Client struct {
	config    Config
	userAgent string
}

// Config The confid structure for the client api for stingray.
type Config struct {
	UserAgent   string
	HTTPTimeout time.Duration
	Host        string
	Port        int
	UserName    string
	Password    string
}
