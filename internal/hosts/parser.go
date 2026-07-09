package hosts

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

func Parse(data []byte) ([]Host, error) {

	text := strings.TrimSpace(string(data))
	if text == "" {
		return []Host{}, nil
	}

	text = "[" + text + "]"

	var hosts []Host

	if err := json.Unmarshal([]byte(text), &hosts); err != nil {
		return nil, err
	}

	return hosts, nil
}

func Marshal(hosts []Host) ([]byte, error) {

	out := make([]string, 0, len(hosts))

	for _, h := range hosts {

		b, err := json.Marshal(h)
		if err != nil {
			return nil, err
		}

		out = append(out, string(b))
	}

	return []byte(strings.Join(out, ",\n")), nil
}

func ReadFile(path string) ([]Host, error) {

	data, err := os.ReadFile(path)
	if err != nil {

		if errors.Is(err, os.ErrNotExist) {
			return []Host{}, nil
		}

		return nil, err
	}

	return Parse(data)
}

func WriteFile(path string, hosts []Host) error {

	data, err := Marshal(hosts)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
