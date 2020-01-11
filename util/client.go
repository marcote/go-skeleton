package util

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

//Client is our vehicle to interact with SWAPI.CO host.
type Client struct {
	BaseURL    *url.URL
	BasePath   string
	httpClient *http.Client
}

//NewClient returns a new HttpClient
func NewClient() *Client {

	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	client := &Client{
		BaseURL: &url.URL{
			Scheme: "http",
			Host:   "swapi.co",
		},
		BasePath: "/api/",
		httpClient: &http.Client{
			Timeout:   time.Second * 10,
			Transport: netTransport,
		},
	}

	return client
}

// NewRequest creates an API request.
func (c *Client) NewRequest(s string) (*http.Request, error) {
	rel, err := url.Parse(c.BasePath + s)
	if err != nil {
		return nil, err
	}

	q := rel.Query()
	q.Set("format", "json")

	rel.RawQuery = q.Encode()

	u := c.BaseURL.ResolveReference(rel)
	fmt.Println("swapi: GET", u.String())

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Testing")
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	// Make sure to close the connection after replying to this request
	req.Close = true

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	if err != nil {
		return nil, fmt.Errorf("error reading response from %s %s: %s", req.Method, req.URL.RequestURI(), err)
	}

	return resp, nil
}
