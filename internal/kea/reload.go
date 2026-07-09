package kea

import "context"

type ReloadRequest struct {
	Command string   `json:"command"`
	Service []string `json:"service"`
}

func (c *Client) Reload() error {

	req := ReloadRequest{
		Command: "config-reload",
		Service: []string{"dhcp4"},
	}

	return c.Call(
		context.Background(),
		req,
		nil,
	)
}
