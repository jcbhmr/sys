#!/usr/bin/env bash
set -Eeuo pipefail
IFS=$'\n\t'

temp_dir=$(mktemp -d)
pushd "$temp_dir"

cleanup() {
    trap - SIGINT SIGTERM ERR EXIT
    cd /
    rm -rf "$temp_dir"
}
trap cleanup SIGINT SIGTERM ERR EXIT

url=$(curl -fsSLo /dev/null -w '%{url_effective}\n' https://github.com/bytecodealliance/wasm-pkg-tools/releases/latest)
tag="${url##*/}"
mkdir -p ~/.local/bin
curl -fsSLo ~/.local/bin/wkg "https://github.com/bytecodealliance/wasm-pkg-tools/releases/download/$tag/wkg-x86_64-unknown-linux-gnu"
chmod +x ~/.local/bin/wkg
