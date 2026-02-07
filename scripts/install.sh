#!/bin/bash
set -e

REPO="TanmayDabhade/orion"

# 1. Detect OS and Arch
OS=$(uname -s)
ARCH=$(uname -m)

if [ "$OS" == "Darwin" ]; then
    OS="Darwin"
elif [ "$OS" == "Linux" ]; then
    OS="Linux"
else
    echo "Unsupported OS: $OS"
    exit 1
fi

if [ "$ARCH" == "x86_64" ]; then
    ARCH="x86_64"
elif [ "$ARCH" == "arm64" ]; then
    ARCH="arm64"
elif [ "$ARCH" == "aarch64" ]; then
     ARCH="arm64"
else
    echo "Unsupported architecture: $ARCH"
    exit 1
fi

# 2. Construct Download URL
# Filename format from .goreleaser.yaml: orion_Darwin_arm64.tar.gz
FILENAME="orion_${OS}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/latest/download/${FILENAME}"

echo "Downloading Orion for ${OS}/${ARCH}..."
echo "  Source: $URL"

# 3. Download and Extract
curl -fsSL "$URL" -o orion.tar.gz
tar -xzf orion.tar.gz orion

# 4. Install
echo "Installing to /usr/local/bin (requires sudo)..."
sudo mv orion /usr/local/bin/o
rm orion.tar.gz

echo "Successfully installed!"
echo "Run 'o --help' to get started."
