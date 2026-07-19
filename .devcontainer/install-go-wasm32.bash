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

git clone https://github.com/Zxilly/go.git Zxilly-go
pushd Zxilly-go
git switch feat/wasm32

pushd src
./make.bash
popd
popd

mkdir -p ~/.local/share
rm -rf ~/.local/share/Zxilly-go-wasm32
rm -rf Zxilly-go/.git
mv Zxilly-go ~/.local/share/Zxilly-go-wasm32
mkdir -p ~/.local/bin
ln -s ~/.local/share/Zxilly-go-wasm32/bin/go ~/.local/bin/Zxilly-go-wasm32-go
ln -s ~/.local/share/Zxilly-go-wasm32/bin/gofmt ~/.local/bin/Zxilly-go-wasm32-gofmt

git clone https://github.com/jellevandenhooff/go.git jellevandenhooff-go
pushd jellevandenhooff-go
git switch wasm32

pushd src
./make.bash
popd
popd

mkdir -p ~/.local/share
rm -rf ~/.local/share/jellevandenhooff-go-wasm32
rm -rf jellevandenhooff-go/.git
mv jellevandenhooff-go ~/.local/share/jellevandenhooff-go-wasm32
mkdir -p ~/.local/bin
ln -s ~/.local/share/jellevandenhooff-go-wasm32/bin/go ~/.local/bin/jellevandenhooff-go-wasm32-go
ln -s ~/.local/share/jellevandenhooff-go-wasm32/bin/gofmt ~/.local/bin/jellevandenhooff-go-wasm32-gofmt
