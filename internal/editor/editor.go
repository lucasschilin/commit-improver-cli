package editor

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Open(path string) error {
	editor := detectEditor()

	cmd := exec.Command(editor, path)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func OpenTempFile() (string, error) {
	file, err := os.CreateTemp("", "cim-cli-commit-*.txt")
	if err != nil {
		return "", err
	}

	path := file.Name()
	file.Close()
	defer os.Remove(path)

	editor := detectEditor()

	editCmd := exec.Command("sh", "-c", editor+" "+path)
	editCmd.Stdin = os.Stdin
	editCmd.Stdout = os.Stdout
	editCmd.Stderr = os.Stderr

	if err := editCmd.Run(); err != nil {
		return "", err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func detectEditor() string {

	// variável do git
	if e := os.Getenv("GIT_EDITOR"); e != "" {
		return e
	}

	// configuração do git
	cmd := exec.Command("git", "config", "--get", "core.editor")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err == nil {
		editor := strings.TrimSpace(out.String())
		if editor != "" {
			return editor
		}
	}

	// VISUAL
	if e := os.Getenv("VISUAL"); e != "" {
		return e
	}

	// EDITOR
	if e := os.Getenv("EDITOR"); e != "" {
		return e
	}

	// fallback por sistema
	switch runtime.GOOS {
	case "windows":
		return "notepad"
	case "darwin":
		return "nano"
	default:
		return "nano"
	}
}
