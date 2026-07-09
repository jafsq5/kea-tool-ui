package ssh

import (
	"bytes"
	"fmt"
	"io"
	"path"
	"strconv"
	"time"
)

func (c *Client) run(cmd string) error {
	session, err := c.Session()
	if err != nil {
		return err
	}
	defer session.Close()

	return session.Run(cmd)
}

func (c *Client) runWithInput(cmd string, r io.Reader) error {
	session, err := c.Session()
	if err != nil {
		return err
	}
	defer session.Close()

	session.Stdin = r

	return session.Run(cmd)
}

func (c *Client) ReadFile(file string) ([]byte, error) {
	session, err := c.Session()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var stdout bytes.Buffer
	session.Stdout = &stdout

	cmd := fmt.Sprintf("cat %s", strconv.Quote(file))

	if err := session.Run(cmd); err != nil {
		return nil, err
	}

	return io.ReadAll(&stdout)
}

func (c *Client) WriteFile(file string, data []byte) error {
	cmd := fmt.Sprintf("cat > %s", strconv.Quote(file))

	return c.runWithInput(cmd, bytes.NewReader(data))
}

func (c *Client) Backup(file string) error {
	backup := fmt.Sprintf(
		"%s.%s.bak",
		file,
		time.Now().Format("20060102-150405"),
	)

	cmd := fmt.Sprintf(
		"cp %s %s",
		strconv.Quote(file),
		strconv.Quote(path.Clean(backup)),
	)

	return c.run(cmd)
}

func (c *Client) Replace(tmpFile, dstFile string) error {
	cmd := fmt.Sprintf(
		"mv %s %s",
		strconv.Quote(tmpFile),
		strconv.Quote(dstFile),
	)

	return c.run(cmd)
}
