package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"
)

// NewClient Creates a new instance of the stingray api client.
func NewClient(config Config) (*Client, error) {
	if _, err := url.Parse(config.HTTPEndpoint); err != nil {
		return nil, err
	}

	return &Client{
		config:        config,
		userAgent:     generateUserAgent(config),
		httpTimeout:   getTimeout(config),
		httpTransport: getTransport(config),
	}, nil
}

func getTransport(config Config) *http.Transport {
	tldConf := &tls.Config{
		InsecureSkipVerify: config.InsecureSkipVerify,
	}

	return &http.Transport{
		TLSClientConfig: tldConf,
	}
}

func getTimeout(config Config) time.Duration {
	if config.HTTPTimeout == 0 {
		return 30 * time.Second
	}
	return config.HTTPTimeout
}

func generateUserAgent(config Config) string {
	if config.UserAgent != "" {
		return fmt.Sprintf(userAgentTemplate, config.UserAgent)
	}

	return fmt.Sprintf(userAgentTemplate, "application name not set")
}

func (client *Client) createRequest(method string, relative string, data []byte) (*http.Request, error) {
	url, err := url.Parse(client.config.HTTPEndpoint)
	if err != nil {
		return nil, err
	}

	url.Path = path.Join(url.Path, relative)

	req, err := http.NewRequest(method, url.String(), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	if len(data) != 0 {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}
	req.Header.Set("accept", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", client.userAgent)
	req.SetBasicAuth(client.config.UserName, client.config.Password)

	return req, nil
}

// GetActionList Default method to fetch action lists from the stingray API.
func (client *Client) GetActionList(url string) (*ActionList, error) {
	request, err := client.createRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{Timeout: client.httpTimeout, Transport: client.httpTransport}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	actionList := ActionList{}

	err = json.NewDecoder(response.Body).Decode(&actionList)
	if err != nil {
		return nil, err
	}

	return &actionList, nil
}
