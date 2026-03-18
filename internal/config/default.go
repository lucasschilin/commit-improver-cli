package config

const DefaultConfigTemplate = `# Commit Improver CLI configuration
#
# Required fields must be filled for the CLI to work.

# Language used for generated commits. Options: 'en' (English), 'pt-BR' (Português do Brasil) or 'es' (Español). Default: 'en'
language: en

# Maximum number of diff lines sent to the LLM. Default: 200
diff_limit: 200

# Maximum number of seconds to wait for a response from the LLM. Default: 20
improvement_request_timeout: 20

# Allow the user to edit the final commit message. Default: false
allow_final_edit: false

# LLM provider used to improve commits. Options: 'gemini'. Default: 'gemini'
provider: gemini

# Model name used by the provider. Default: 'gemini-2.5-flash'
model: gemini-2.5-flash

gemini:
  # Gemini API key
  api_key: # REQUIRED (not indicated for --repo configurations, as it will be versioned with the code)

`
