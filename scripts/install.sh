#!/bin/bash

# Create the logs and config directories
mkdir -p $HOME/.routnd/{logs,config}

# Map architecture to binary name
ARCH=$(uname -m)
case $ARCH in
    i386|i686) ARCH=x86 ;;
    armv6*) ARCH=armv6 ;;
    armv7*) ARCH=armv7 ;;
    aarch64*) ARCH=arm64 ;;
esac

# Download the binary
GITHUB_LATEST_VERSION=$(curl -L -s -H 'Accept: application/json' https://github.com/deestarks/routnd/releases/latest | sed -e 's/.*"tag_name":"\([^"]*\)".*/\1/')
GITHUB_FILE="routnd_${GITHUB_LATEST_VERSION//v/}_$(uname -s)_${ARCH}.tar.gz"

curl -L -o routnd.tar.gz "https://github.com/deestarks/routnd/releases/download/${GITHUB_LATEST_VERSION}/${GITHUB_FILE}"
tar xzvf routnd.tar.gz routnd
install -Dm 755 routnd -t $HOME/.local/bin
rm routnd routnd.tar.gz