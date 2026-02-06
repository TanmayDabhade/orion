#!/bin/bash
set -e

REPO="user/orion" # Replace with actual repo
VERSION="latest"
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [ "$ARCH" == "x86_64" ]; then
    ARCH="amd64"
elif [ "$ARCH" == "arm64" ]; then
    ARCH="arm64"
else
    echo "Unsupported architecture: $ARCH"
    exit 1
fi

BINARY="orion_${OS}_${ARCH}"
URL="https://github.com/${REPO}/releases/${VERSION}/download/${BINARY}"

echo "Downloading Orion..."
# In a real scenario, we would download here. For now, we simulate.
# curl -L -o orion $URL

echo "Simulated download from: $URL"
echo "To install, move 'orion' to /usr/local/bin"
