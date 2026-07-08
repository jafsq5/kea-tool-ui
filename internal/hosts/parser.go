package hosts

import (
	"encoding/json"
	"os"
	"strings"
)

func ReadFile(path string) ([]Host, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	text := strings.TrimSpace(string(data))

	if text == "" {
		return []Host{}, nil
	}

	text = "[" + text + "]"

	var hosts []Host

	err = json.Unmarshal([]byte(text), &hosts)
	if err != nil {
		return nil, err
	}

	return hosts, nil
}

func WriteFile(path string, hosts []Host) error {

	lines := make([]string, 0, len(hosts))

	for _, h := range hosts {

		b, err := json.Marshal(h)
		if err != nil {
			return err
		}

		lines = append(lines, string(b))
	}

	text := strings.Join(lines, ",\n")

	return os.WriteFile(path, []byte(text), 0644)
}
