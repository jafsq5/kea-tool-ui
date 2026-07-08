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

	return hosts, err
}

func WriteFile(path string, hosts []Host) error {

	out := make([]string, 0, len(hosts))

	for _, h := range hosts {

		b, err := json.Marshal(h)
		if err != nil {
			return err
		}

		out = append(out, string(b))
	}

	return os.WriteFile(
		path,
		[]byte(strings.Join(out, ",\n")),
		0644,
	)
}
