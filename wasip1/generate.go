//go:build generate

//go:generate env -C generate cargo build
//go:generate ./generate/target/debug/generate

package wasip1
