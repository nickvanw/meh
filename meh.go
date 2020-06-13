package meh

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://api.meh.com/"
	apiVer  = "1"
)

// Client is a Meh Go API client
type Client struct {
	key string
	url *url.URL

	client *http.Client
}

// NewClient returns a new Meh API Client
func NewClient(conf ...func(c *Client)) *Client {
	base, _ := url.Parse(baseURL)

	c := &Client{
		url:    base,
		client: http.DefaultClient,
	}
	for _, v := range conf {
		v(c)
	}
	return c
}

// Current returns the current deal on the meh.com homepage
func (c *Client) Current() (*Current, error) {
	u := "current.json"

	req, err := c.newRequest("GET", u)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Set("apikey", c.key)
	req.URL.RawQuery = q.Encode()

	cur := new(Current)

	_, err = c.do(req, cur)
	if err != nil {
		return nil, err
	}

	return cur, nil
}

func (c *Client) newRequest(method, urlStr string) (*http.Request, error) {
	rel, err := url.Parse(fmt.Sprintf("%s/%s", apiVer, urlStr))
	if err != nil {
		return nil, err
	}

	u := c.url.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if v == nil {
		return resp, nil
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return nil, err
	}

	return resp, nil
}
