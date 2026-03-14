package prompt

import "fmt"

func Build(message string, diff string, language string) string {
	return fmt.Sprintf(`
		You are a senior software engineer.

		Rewrite the following commit message using Conventional Commits.

		Rules:
		- Translate to %s
		- Maximum 72 characters
		- Use imperative mood
		- Return only the commit message

		Original message:
		%s

		Git diff:
		%s`,
		language, message, diff)
}
