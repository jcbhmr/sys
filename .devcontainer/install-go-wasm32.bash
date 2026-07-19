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
mv go ~/.local/share/Zxilly-go-wasm32
mkdir -p ~/.local/bin
echo '#!/usr/bin/env bash' > ~/.local/bin/Zxilly-go-wasm32
echo 'export GOARCH=wasm32' >> ~/.local/bin/Zxilly-go-wasm32
echo 'exec ~/.local/share/Zxilly-go-wasm32/bin/go "$@"' >> ~/.local/bin/Zxilly-go-wasm32
chmod +x ~/.local/bin/Zxilly-go-wasm32

git clone https://github.com/jellevandenhooff/go.git jellevandenhooff-go
pushd jellevandenhooff-go
git switch wasm32

pushd src
./make.bash
popd
popd

mkdir -p ~/.local/share
rm -rf ~/.local/share/jellevandenhooff-go-wasm32
mv go ~/.local/share/jellevandenhooff-go-wasm32
mkdir -p ~/.local/bin
echo '#!/usr/bin/env bash' > ~/.local/bin/jellevandenhooff-go-wasm32
echo 'export GOARCH=wasm32' >> ~/.local/bin/jellevandenhooff-go-wasm32
echo 'exec ~/.local/share/jellevandenhooff-go-wasm32/bin/go "$@"' >> ~/.local/bin/jellevandenhooff-go-wasm32
chmod +x ~/.local/bin/jellevandenhooff-go-wasm32
