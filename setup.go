package meh

import (
	"net/http"
	"net/url"
)

func WithKey(key string) func(c *Client) {
	return func(c *Client) {
		c.key = key
	}
}

func WithURL(urlStr string) func(c *Client) {
	u, _ := url.Parse(urlStr)
	return func(c *Client) {
		c.url = u
	}
}

func WithClient(client *http.Client) func(c *Client) {
	return func(c *Client) {
		c.client = client
	}
}
