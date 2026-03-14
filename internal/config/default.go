package config

const DefaultConfigTemplate = `# Commit Improver CLI configuration
#
# Required fields must be filled for the CLI to work.

# LLM provider used to improve commits. Options: 'gemini'. Default: 'gemini'
provider: gemini #REQUIRED

# Model name used by the provider. Default: 'gemini-2.5-flash'
model: gemini-2.5-flash #REQUIRED

# Language used for generated commits. Options: 'en' (English), 'pt-BR' (Português do Brasil) or 'es' (Español). Default: 'en'
language: en #REQUIRED

# Maximum number of diff lines sent to the LLM. Default: 200
diff_limit: 200 #REQUIRED

gemini:
  # Gemini API key
  api_key: # REQUIRED (not indicated for --repo configurations, as it will be versioned with the code)

`
