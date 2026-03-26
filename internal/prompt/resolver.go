package prompt

func Resolve(repoRoot string) (string, bool, error) { // TODO: receive here diff and message
	if repoRoot != "" {
		if content, ok, err := Load(RepoSharedPromptPath(repoRoot)); err != nil {
			return "", false, err
		} else if ok {
			return content, true, nil
		}
	}

	if repoRoot != "" {
		if content, ok, err := Load(RepoUserPromptPath(repoRoot)); err != nil {
			return "", false, err
		} else if ok {
			return content, true, nil
		}
	}

	if content, ok, err := Load(GlobalPromptPath()); err != nil {
		return "", false, err
	} else if ok {
		return content, true, nil
	}

	return "", false, nil // TODO: fallback is prompt.Build
}
