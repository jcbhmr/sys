## Development

- https://github.com/bytecodealliance/wasi-rs/tree/4b007f5b57b23a254e6fc0b92de6d6431371cb76/crates/wasip1
- https://github.com/golang/go/issues/63131

```sh
GOOS=wasip1 GOARCH=wasm32 Zxilly-go-wasm32-go test -c
wasmtime wasip1.test
```

```sh
# Doesn't work.
GOOS=wasip1 GOARCH=wasm32 jellevandenhooff-go-wasm32-go test -c
wasmtime wasip1.test
```
