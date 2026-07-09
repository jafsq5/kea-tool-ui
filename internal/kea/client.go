package kea

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	url    string
	client *http.Client
}

func New(url string) *Client {
	return &Client{
		url: url,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) Call(ctx context.Context, req any, resp any) error {

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.url,
		bytes.NewReader(body),
	)
	if err != nil {
		return err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("kea returned %s", httpResp.Status)
	}

	if resp == nil {
		return nil
	}

	return json.NewDecoder(httpResp.Body).Decode(resp)
}
