package stingray

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

func (client *Client) putResource(url string, object interface{}) error {
	data, err := json.Marshal(&object)

	request, err := client.createRequest("PUT", url, data)
	if err != nil {
		return err
	}

	httpClient := &http.Client{Timeout: client.httpTimeout, Transport: client.httpTransport}

	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Error status from the stingray api status code: %s", response.Status)
	}

	// stingray api returns the changed object right now we ignore this.
	/*	defer response.Body.Close()

		err = json.NewDecoder(response.Body).Decode(&object)
		if err != nil {
			return err
		}*/

	return nil
}

func (client *Client) getResource(url string, object interface{}) error {
	request, err := client.createRequest("GET", url, nil)
	if err != nil {
		return err
	}

	httpClient := &http.Client{Timeout: client.httpTimeout, Transport: client.httpTransport}

	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&object)
	if err != nil {
		return err
	}

	return nil
}

// GetTrafficIPGroup Get a specific traffic IP group resource from the stingray API.
func (client *Client) GetTrafficIPGroup(url string) (*TrafficIPGroup, error) {
	trafficIPGroup := TrafficIPGroup{}

	err := client.getResource(url, &trafficIPGroup)
	if err != nil {
		return nil, err
	}

	return &trafficIPGroup, nil
}

// SetTrafficIPGroup Sets a specific traffic IP group resource from the stingray API.
func (client *Client) SetTrafficIPGroup(url string, trafficIPGroup TrafficIPGroup) error {
	err := client.putResource(url, &trafficIPGroup)
	if err != nil {
		return err
	}

	return nil
}

// GetPool Get a specific pool resource from the stingray API.
func (client *Client) GetPool(url string) (*Pool, error) {
	pool := Pool{}

	err := client.getResource(url, &pool)
	if err != nil {
		return nil, err
	}

	return &pool, nil
}

// SetPool Set a specific pool resource from the stingray API.
func (client *Client) SetPool(url string, pool Pool) error {
	err := client.putResource(url, &pool)
	if err != nil {
		return err
	}

	return nil
}

// GetVirtualServer Get a specific virtual server resource from the stingray API.
func (client *Client) GetVirtualServer(url string) (*VirtualServer, error) {
	virtualServer := VirtualServer{}

	err := client.getResource(url, &virtualServer)
	if err != nil {
		return nil, err
	}

	return &virtualServer, nil
}

// SetVirtualServer Set a specific virtual server resource from the stingray API.
func (client *Client) SetVirtualServer(url string, virtualServer VirtualServer) error {
	err := client.putResource(url, &virtualServer)
	if err != nil {
		return err
	}

	return nil
}
