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

### Manual Install
Download the binary for your system from the [Releases page](https://github.com/TanmayDabhade/orion/releases).

## Configuration

Orion config is stored at `~/.config/orion/config.yaml`.

### Enable Google Gemini AI (Recommended)
By default, Orion tries to use Ollama (local). To use Google's Gemini models:

1.  Get an API Key from [Google AI Studio](https://aistudio.google.com/).
2.  Update your config:

```yaml
ai_provider: gemini
ai_key: YOUR_GEMINI_API_KEY_HERE
ai_model: gemini-pro
search_engine: https://google.com/search?q=%s
risk_threshold: medium
features:
  ai_fallback: true
```

## Usage

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

### App Launching
Orion smart-matches `.app` names:
```bash
o outlook      # Launches Outlook.app
o spotify      # Launches Spotify.app
```

### AI Commands (Natural Language)
If no shortcut matches, Orion routes to AI (if configured):
```bash
o search discrete math cheat sheet
# -> Opens Google Search for "discrete math cheat sheet"

o open youtube
# -> Opens youtube.com
```

### Management
```bash
o list         # List all shortcuts
o doctor       # Check system health & config
o update       # Check for updates
```

## Contributing
1.  Fork the repo.
2.  Feature branch: `git checkout -b feature/xyz`
3.  Commit & Push.
4.  Submit PR.

## License
MIT
# orion
