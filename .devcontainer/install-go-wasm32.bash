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

git clone https://github.com/Zxilly/go.git
pushd go
git switch feat/wasm32

pushd src
./make.bash
popd
popd

mkdir -p ~/.local/share
rm -rf ~/.local/share/go-wasm32
rm -rf go/.git
mv go ~/.local/share/go-wasm32
mkdir -p ~/.local/bin
ln -s ~/.local/share/go-wasm32/bin/go ~/.local/bin/gowasm32
