package prompt

import (
	"fmt"
	"strings"
)

func ResolveTemplate(repoRoot string) (string, error) {
	prompt := Default()

	content, ok, err := Load(GlobalPromptPath())
	if err != nil {
		return "", err
	} else if ok == true && content != "" {
		prompt = content
	}

	content, ok, err = Load(RepoUserPromptPath(repoRoot))
	if err != nil {
		return "", err
	} else if ok == true && content != "" {
		prompt = content
	}

	content, ok, err = Load(RepoSharedPromptPath(repoRoot))
	if err != nil {
		return "", err
	} else if ok == true && content != "" {
		prompt = content
	}

	return prompt, nil
}

func applyReplaces(prompt string, message string, diff string, language string) string {
	replaces := map[string]string{
		"MESSAGE":  message,
		"DIFF":     diff,
		"LANGUAGE": language,
	}

	for key, value := range replaces {
		prompt = strings.ReplaceAll(prompt, fmt.Sprintf("{{%s}}", key), value)
	}

	return prompt
}

func Prompt(repoRoot string, message string, diff string, language string) (string, error) {
	prompt, err := ResolveTemplate(repoRoot)
	if err != nil {
		return "", err
	}

	return applyReplaces(prompt, message, diff, language), nil
}
