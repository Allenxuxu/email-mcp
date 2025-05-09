#!/bin/bash
set -e -u

REPO="Allenxuxu/email-mcp"
BINARY_NAME="send-email"

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
case "$ARCH" in
  x86_64|amd64)
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

API_URL="https://api.github.com/repos/${REPO}/releases/latest"
DOWNLOAD_BASE_URL="https://github.com/${REPO}/releases/download"
TAG_NAME=$(curl -s $API_URL | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
echo "TAG_NAME: ${TAG_NAME}"
if [ -z "$TAG_NAME" ]; then
  echo "Failed to get latest release version from GitHub: ${API_URL}"
  exit 1
fi

if [ "$OS" = "darwin" ]; then
  FILENAME="${BINARY_NAME}-${TAG_NAME}-darwin-${ARCH}.tar.gz"
elif [ "$OS" = "linux" ]; then
  FILENAME="${BINARY_NAME}-${TAG_NAME}-linux-${ARCH}.tar.gz"
elif [ "$OS" = "windows" ] || [[ "$OS" == *mingw* ]]; then
  FILENAME="${BINARY_NAME}-${TAG_NAME}-windows-${ARCH}.zip"
else
  echo "Unsupported OS: $OS"
  exit 1
fi

API_URL="https://api.github.com/repos/${REPO}/releases/latest"
DOWNLOAD_BASE_URL="https://github.com/${REPO}/releases/download"
FILE_URL="${DOWNLOAD_BASE_URL}/${TAG_NAME}/${FILENAME}"

echo "Detected platform: ${OS}-${ARCH}"
echo "Downloading: ${FILE_URL}"

curl -LO "${FILE_URL}"

if [[ "$FILENAME" == *.tar.gz ]]; then
  tar -xzf "$FILENAME"
elif [[ "$FILENAME" == *.zip ]]; then
  unzip "$FILENAME"
fi

rm "$FILENAME"

echo "Download and extraction completed successfully."

cat <<-'EOM'

Send Email MCP has been downloaded to the current directory.
You can run:

./send-email -h

EOM
