#!/bin/bash
set -e

echo "Installing MaxGPT..."

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    aarch64|arm64)
        ARCH="arm64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

BINARY_NAME="maxgpt-${OS}-${ARCH}"
DOWNLOAD_URL="https://github.com/yosiwizman/local-maxgpt/releases/latest/download/${BINARY_NAME}"

echo "Downloading MaxGPT for ${OS}/${ARCH}..."
curl -L -o maxgpt "${DOWNLOAD_URL}"
chmod +x maxgpt

if [ -w "/usr/local/bin" ]; then
    mv maxgpt /usr/local/bin/
    echo "MaxGPT installed to /usr/local/bin/maxgpt"
else
    echo "Installing to ~/.local/bin/ (add to PATH if needed)"
    mkdir -p ~/.local/bin
    mv maxgpt ~/.local/bin/
    echo "MaxGPT installed to ~/.local/bin/maxgpt"
fi

echo "MaxGPT installation complete!"
echo "Run 'maxgpt --help' to get started."
