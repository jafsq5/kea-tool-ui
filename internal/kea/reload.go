package kea

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	url string
	http *http.Client
}

func New(url string) *Client {
	return &Client{
		url: url,
		http: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) Reload() error {

	req := map[string]any{
		"command": "config-reload",
		"service": []string{"dhcp4"},
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := c.http.Post(
		c.url,
		"application/json",
		bytes.NewReader(body),
	)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("kea returned %s", resp.Status)
	}

	return nil
}
