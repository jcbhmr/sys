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

url=$(curl -fsSLo /dev/null -w '%{url_effective}\n' https://github.com/bytecodealliance/wasm-tools/releases/latest)
tag="${url##*/}"
version="${tag#v}"
curl -fsSLo wasm-tools.tar.gz "https://github.com/bytecodealliance/wasm-tools/releases/download/$tag/wasm-tools-$version-x86_64-linux.tar.gz"
tar -xzf wasm-tools.tar.gz
mkdir -p ~/.local/share
rm -rf ~/.local/share/wasm-tools
mv ./wasm-tools-* ~/.local/share/wasm-tools
mkdir -p ~/.local/bin
ln -s ~/.local/share/wasm-tools/wasm-tools ~/.local/bin/wasm-tools
