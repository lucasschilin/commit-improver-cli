<p align="center">
  
  <img src="docs/images/logo.png" width="200" />

</p>
<p align="center">
  <img src="https://img.shields.io/badge/platform-linux%20%7C%20macOS%20%7C%20windows-blue" />
  <img src="https://img.shields.io/github/v/release/lucasschilin/cim-cli" />
  <img src="https://img.shields.io/github/release-date/lucasschilin/cim-cli" />
  <img src="https://img.shields.io/github/stars/lucasschilin/cim-cli?style=flat-square" />
  </p>
<p align="center">
  <img src="https://img.shields.io/github/go-mod/go-version/lucasschilin/cim-cli" />
  <img src="https://img.shields.io/badge/built%20with-cobra-00ADD8" />
</p>


> Never write a bad commit message again.

Use AI to automatically transform your commits into clear, standardized Conventional Commits — with or without an initial message.

⭐ If this project helps you, consider giving it a star.

> ![alt text](docs/images/hook.gif)

---

## Why?

Writing good commit messages is harder than it should be.

- You write things like "fix", "adjustments", or "wip"
- Your team uses different styles
- You know Conventional Commits… but don’t always follow them
- You don’t want to think too much about it

## What is cim-cli?

cim-cli acts like a commit copilot.
It intercepts your commit message, analyzes your changes, and suggests a better version using AI.

- Follows Conventional Commits
- Understands your code via git diff
- Works in any repository
- Lets you review before applying

You can even skip writing a message entirely.
cim-cli can generate a meaningful commit message based only on your changes.

---

## Installation
You can download prebuilt binaries for Linux, macOS and Windows.

Download the latest release from GitHub and install it in a few steps.

### Linux / macOS

```bash
# Download the latest version
curl -L https://github.com/lucasschilin/cim-cli/releases/latest/download/cim-cli_Linux_x86_64.tar.gz -o cim-cli.tar.gz
# Extract the binary
tar -xzf cim-cli.tar.gz
# Move to a directory in your PATH
sudo mv cim-cli /usr/local/bin/
# Make sure it is executable
chmod +x /usr/local/bin/cim-cli
```

### Windows (PowerShell)
```PowerShell
# Download the latest version
Invoke-WebRequest -Uri https://github.com/lucasschilin/cim-cli/releases/latest/download/cim-cli_Windows_x86_64.zip -OutFile cim-cli.zip
# Extract the zip
Expand-Archive cim-cli.zip
# Move the binary to a folder in your PATH (example)
Move-Item .\cim-cli.exe "$env:USERPROFILE\AppData\Local\Microsoft\WindowsApps\"
```

### Verify Installation
```bash
cim-cli --version
```

---

## Configuration

cim-cli provides a flexible configuration system that works across different scopes.

### Configuration levels

You can configure cim-cli in three levels:

- **Global** → applies to all repositories (`~/.cim-cli/`)
- **Repo-user** → applies only to you, inside a specific repository (`.git/.cim-cli/`)
- **Repo** → shared with the team and versioned in the repository (`.cim-cli/`)

### Priority

Configuration is merged by priority:
```
global < repo-user < repo
```

The repository configuration always has the highest priority, ensuring consistency across teams.

### Smart merging

Each configuration field is resolved independently.

For example:

- Language can come from `repo`
- Diff limit can come from `global`

This allows flexible setups without overriding everything.

### Examples

#### Different language per repository

You can use English globally:

```bash
cim-cli config params edit --global
```
And Portuguese in a specific repository:
```bash
cim-cli config params edit --repo-user
```

You can check the configurations options using:
```bash
cim-cli config template
``` 

### Commands

Edit configuration:
```bash
cim-cli config params edit --global
cim-cli config params edit --repo
cim-cli config params edit --repo-user
```
View configuration:
```bash
cim-cli config params show --global
cim-cli config params show --repo
cim-cli config params show --repo-user
```
View the final merged configuration:
```bash
cim-cli config params show
```

### API Key Configuration
You need to set up an API key configuration in the configuration edit

---

## Setup (Git Hook)

To enable automatic commit message improvement, install the git hook:

```bash
# Into a git repository
cim-cli init
```
Now, every time you run:

`git commit`

cim-cli will intercept and improve your commit message automatically.

## Usage

### Automatic (Git Hook)

Just commit as usual:

```bash
# example
git commit -m "fix"
```
cim-cli will suggest an improved message before completing the commit.

### Manual

You can also improve messages manually:
```bash
cim-cli improve -m "fix login"
```
or:
```bash
cim-cli improve
```
> ![alt text](docs/images/improve.gif)

This is useful when you want to refine a message outside the commit flow.

---

## Why cim-cli?

There are many tools that help enforce commit message standards, like commit linters and git hooks.

cim-cli takes a different approach.

Instead of rejecting your commits, it helps you write better ones.

- It **improves** your message instead of validating it
- It uses **AI + git diff** to understand your changes
- It works **globally**, not tied to a single repository
- It keeps you in control with **review and optional editing**

Most tools focus on enforcing rules.

cim-cli focuses on making it effortless to follow them.

---

## Support the project

If this project helps you:

⭐ Consider giving it a star  
🤝 Contributions are welcome