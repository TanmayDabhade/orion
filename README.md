# Orion (CLI: `o`)

**Orion** is a natural-language terminal assistant that converts conversational user commands into safe, structured system automation. It eliminates repetitive navigation by acting as a smart launchpad, shortcut manager, and AI-powered command router.

## Features

- **Smart Shortcuts**: Launch apps or URLs with short aliases (`o d2l` → `chrome https://d2l.msu.edu`).
- **App Launcher**: Automatically detects apps (`o outlook` → `open -a Outlook.app`).
- **AI Router**: Intelligently handles natural language requests using **Google Gemini** or **Ollama** (e.g., "search for discrete math cheat sheet" -> opens browser search).
- **Safety First**: Conservative risk gating ensures no dangerous commands run without confirmation.
- **Cross-Platform**: Built in Go for macOS (Phase 1), with Linux/Windows support planned.

## Installation

### Automatic Install (macOS & Linux)
```bash
curl -fsSL https://github.com/TanmayDabhade/orion/releases/latest/download/install.sh | sh
```

## Configuration

Orion config is stored at `~/.config/orion/config.yaml`.

### 1. Enable Google Gemini AI (Recommended)
By default, Orion tries to use Ollama (local). To use Google's Gemini models:

1.  Get an API Key from [Google AI Studio](https://aistudio.google.com/).
2.  Run `o doctor` to verify config location.
3.  Update your config:

```yaml
ai_provider: gemini
ai_key: YOUR_GEMINI_API_KEY_HERE
ai_model: gemini-pro
features:
  ai_fallback: true
```

### 2. Enable Shell Completion
Orion supports autocompletion for Bash, Zsh, Fish, and PowerShell.
Run the following to see setup instructions for your shell:

```bash
o completion --help
```

**Example (Zsh):**
```bash
# Add to ~/.zshrc
echo "autoload -U compinit; compinit" >> ~/.zshrc
o completion zsh > "${fpath[1]}/_o"
```

## Usage

### App Launching (Zero Config)
Orion automatically detects installed applications on macOS.

```bash
o slack        # Opens Slack.app
o chrome       # Opens Google Chrome.app
o code .       # Opens current directory in VS Code
```

See all detected apps:
```bash
o list --apps
```

### Shortcuts
Add frequently used URLs or commands:
```bash
o add d2l "open https://d2l.msu.edu"
o add mail "open https://mail.google.com"
```
Use them:
```bash
o d2l
o mail
```

### AI Commands (Natural Language)
If no shortcut or app matches, Orion routes to AI:
```bash
o search discrete math cheat sheet
# -> Opens Google Search for "discrete math cheat sheet"

o how do i check open ports?
# -> Suggests: lsof -i -P
```

### Management
```bash
o list         # List all shortcuts
o list --apps  # List detected applications
o doctor       # Check system health & config
o update       # Update Orion to the latest version
```

## Contributing
1.  Fork the repo.
2.  Feature branch: `git checkout -b feature/xyz`
3.  Commit & Push.
4.  Submit PR.

## License
MIT
