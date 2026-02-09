#!/bin/bash
set -e

REPO="TanmayDabhade/orion"
CONFIG_DIR="$HOME/.config/orion"
CONFIG_FILE="$CONFIG_DIR/config.yaml"

echo "════════════════════════════════════════════════════════════════"
echo "                    🚀 Orion Installation                       "
echo "════════════════════════════════════════════════════════════════"
echo ""

# 1. Detect OS and Arch
OS=$(uname -s)
ARCH=$(uname -m)

if [ "$OS" == "Darwin" ]; then
    OS="Darwin"
elif [ "$OS" == "Linux" ]; then
    OS="Linux"
else
    echo "❌ Unsupported OS: $OS"
    exit 1
fi

if [ "$ARCH" == "x86_64" ]; then
    ARCH="x86_64"
elif [ "$ARCH" == "arm64" ]; then
    ARCH="arm64"
elif [ "$ARCH" == "aarch64" ]; then
    ARCH="arm64"
else
    echo "❌ Unsupported architecture: $ARCH"
    exit 1
fi

echo "📦 Detected: ${OS}/${ARCH}"
echo ""

# 2. Check for Ollama
echo "🔍 Checking for Ollama..."
OLLAMA_AVAILABLE=false
if curl -s --connect-timeout 2 http://localhost:11434/api/tags > /dev/null 2>&1; then
    echo "   ✅ Ollama is running at localhost:11434"
    OLLAMA_AVAILABLE=true
else
    echo "   ⚠️  Ollama not detected (not running or not installed)"
    echo "   💡 Install Ollama from https://ollama.ai for local AI processing"
fi
echo ""

# 3. Prompt for Gemini API Key
echo "🔑 Gemini API Key Setup"
echo "   Orion can use Google's Gemini API for AI features."
echo "   Get a free API key at: https://aistudio.google.com/apikey"
echo ""
read -p "   Enter your Gemini API Key (leave blank to skip): " GEMINI_API_KEY
echo ""

# 4. Determine AI Provider
if [ -n "$GEMINI_API_KEY" ]; then
    AI_PROVIDER="gemini"
    AI_MODEL="gemini-1.5-flash"
    echo "   ✅ Will use Gemini as AI provider"
elif [ "$OLLAMA_AVAILABLE" = true ]; then
    AI_PROVIDER="ollama"
    AI_MODEL="llama3.1"
    echo "   ✅ Will use Ollama as AI provider"
else
    AI_PROVIDER="ollama"
    AI_MODEL="llama3.1"
    echo "   ⚠️  No AI provider configured. Install Ollama or add Gemini key later."
fi
echo ""

# 5. Create config directory and file
echo "📁 Creating configuration..."
mkdir -p "$CONFIG_DIR"

cat > "$CONFIG_FILE" << EOF
# Orion Configuration
# Generated during installation on $(date)

ai_provider: $AI_PROVIDER
ai_model: $AI_MODEL
ai_endpoint: http://localhost:11434
ai_key: $GEMINI_API_KEY
search_engine: https://google.com/search?q=%s
risk_threshold: medium
features:
  ai_fallback: true
EOF

echo "   ✅ Config saved to $CONFIG_FILE"
echo ""

# 6. Construct Download URL
FILENAME="orion_${OS}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/latest/download/${FILENAME}"

echo "📥 Downloading Orion..."
echo "   Source: $URL"

# 7. Download and Extract
curl -fsSL "$URL" -o orion.tar.gz
tar -xzf orion.tar.gz orion

# 8. Install binary
echo ""
echo "🔧 Installing to /usr/local/bin (requires sudo)..."
sudo mv orion /usr/local/bin/o
rm orion.tar.gz

# 9. Run setup to index applications
echo ""
echo "📱 Indexing installed applications..."
/usr/local/bin/o setup

echo ""
echo "════════════════════════════════════════════════════════════════"
echo "                    ✅ Installation Complete!                   "
echo "════════════════════════════════════════════════════════════════"
echo ""
echo "🎉 Orion is ready to use!"
echo ""
echo "   Quick start:"
echo "     o open safari           - Open an application"
echo "     o github.com            - Open a website"
echo "     o search golang tips    - Search the web"
echo ""
echo "   Commands:"
echo "     o --help                - Show all options"
echo "     o list                  - Show all shortcuts"
echo "     o setup                 - Re-index applications"
echo ""
