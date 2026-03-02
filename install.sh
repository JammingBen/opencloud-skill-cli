#!/bin/bash

set -e

# Configuration
REPO="JammingBen/opencloud-skill-cli"
BINARY_NAME="oc-cli"
INSTALL_DIR="/usr/local/bin"

# Detect OS
OS_NAME=$(uname -s)
case "${OS_NAME}" in
    Darwin*)  OS="Darwin" ;;
    Linux*)   OS="Linux" ;;
    *)        echo "Error: Unsupported OS ${OS_NAME}. Only Darwin and Linux are supported." && exit 1 ;;
esac

# Detect Architecture
ARCH_NAME=$(uname -m)
case "${ARCH_NAME}" in
    x86_64*)  ARCH="x86_64" ;;
    arm64*|aarch64*)   ARCH="arm64" ;;
    *)        echo "Error: Unsupported architecture ${ARCH_NAME}." && exit 1 ;;
esac

# Construct download URL for the latest release
DOWNLOAD_URL="https://github.com/${REPO}/releases/latest/download/${BINARY_NAME}_${OS}_${ARCH}.tar.gz"

echo "Downloading ${BINARY_NAME} for ${OS} ${ARCH}..."
TEMP_DIR=$(mktemp -d)
curl -sSL "${DOWNLOAD_URL}" -o "${TEMP_DIR}/${BINARY_NAME}.tar.gz"

echo "Extracting..."
tar -xzf "${TEMP_DIR}/${BINARY_NAME}.tar.gz" -C "${TEMP_DIR}"

echo "Installing to ${INSTALL_DIR}..."
if [ -w "${INSTALL_DIR}" ]; then
    mv "${TEMP_DIR}/${BINARY_NAME}" "${INSTALL_DIR}/"
else
    echo "Requesting sudo for installation to ${INSTALL_DIR}..."
    sudo mv "${TEMP_DIR}/${BINARY_NAME}" "${INSTALL_DIR}/"
fi

chmod +x "${INSTALL_DIR}/${BINARY_NAME}"
rm -rf "${TEMP_DIR}"

echo "Successfully installed ${BINARY_NAME} to ${INSTALL_DIR}"
${BINARY_NAME} version
