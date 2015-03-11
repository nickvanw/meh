package meh

import (
	"net/http"
	"net/url"
)

// WithKey is a config method to set up the meh.com API key
func WithKey(key string) func(c *Client) {
	return func(c *Client) {
		c.key = key
	}
}

// WithURL changes the base URL of the meh.com API
func WithURL(urlStr string) func(c *Client) {
	u, _ := url.Parse(urlStr)
	return func(c *Client) {
		c.url = u
	}
}

// WithClient changes the http client from the default API client
func WithClient(client *http.Client) func(c *Client) {
	return func(c *Client) {
		c.client = client
	}
}
