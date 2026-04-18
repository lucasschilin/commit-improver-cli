package commit

import (
	"os"
	"strings"
)

func ReadCommitMessage(path string) (string, bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", false, err
	}

	raw := string(data)
	lines := strings.Split(raw, "\n")

	if len(lines) > 0 && strings.TrimSpace(lines[0]) == "#ignore" {
		return raw, true, nil
	}

	var filtered []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "#") {
			continue
		}

		filtered = append(filtered, trimmed)
	}

	message := strings.Join(filtered, "\n")
	message = strings.TrimSpace(message)

	return message, false, nil
}

func WriteCommitMessage(path string, message string) error {

	data := []byte(message)

	return os.WriteFile(path, data, 0644)
}
