package commit

import (
	"os"
	"strings"
)

func ReadCommitMessage(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	raw := string(data)

	var filtered []string

	lines := strings.Split(raw, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "#") {
			continue
		}

		filtered = append(filtered, trimmed)
	}

	message := strings.Join(filtered, "\n")
	message = strings.TrimSpace(message)

	return message, nil
}

func WriteCommitMessage(path string, message string) error {

	data := []byte(message)

	return os.WriteFile(path, data, 0644)
}
