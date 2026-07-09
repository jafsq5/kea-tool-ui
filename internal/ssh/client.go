package ssh

import (
	"fmt"
	"net"
	"os"
	"time"

	gossh "golang.org/x/crypto/ssh"
)

type Client struct {
	Host       string
	Port       int
	User       string
	PrivateKey string
	Timeout    time.Duration

	config *gossh.ClientConfig
	client *gossh.Client
}

func New(host string, port int, user, privateKey string, timeout time.Duration) (*Client, error) {

	key, err := os.ReadFile(privateKey)
	if err != nil {
		return nil, err
	}

	signer, err := gossh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	cfg := &gossh.ClientConfig{
		User: user,
		Auth: []gossh.AuthMethod{
			gossh.PublicKeys(signer),
		},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
		Timeout:         timeout,
	}

	return &Client{
		Host:       host,
		Port:       port,
		User:       user,
		PrivateKey: privateKey,
		Timeout:    timeout,
		config:     cfg,
	}, nil
}

func (c *Client) Connect() error {

	if c.client != nil {
		return nil
	}

	addr := net.JoinHostPort(c.Host, fmt.Sprintf("%d", c.Port))

	client, err := gossh.Dial("tcp", addr, c.config)
	if err != nil {
		return err
	}

	c.client = client

	return nil
}

func (c *Client) Session() (*gossh.Session, error) {

	if err := c.Connect(); err != nil {
		return nil, err
	}

	return c.client.NewSession()
}

func (c *Client) Close() error {

	if c.client == nil {
		return nil
	}

	err := c.client.Close()
	c.client = nil

	return err
}
