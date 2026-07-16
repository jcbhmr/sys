//go:build wasip2

package wasip2_test

import (
	"testing"

	_ "go.jcbhmr.com/sys/wasip2/cli/environment"
	_ "go.jcbhmr.com/sys/wasip2/cli/exit"
	_ "go.jcbhmr.com/sys/wasip2/cli/stderr"
	_ "go.jcbhmr.com/sys/wasip2/cli/stdin"
	_ "go.jcbhmr.com/sys/wasip2/cli/stdout"
	_ "go.jcbhmr.com/sys/wasip2/cli/terminal-input"
	_ "go.jcbhmr.com/sys/wasip2/cli/terminal-output"
	_ "go.jcbhmr.com/sys/wasip2/cli/terminal-stderr"
	_ "go.jcbhmr.com/sys/wasip2/cli/terminal-stdin"
	_ "go.jcbhmr.com/sys/wasip2/cli/terminal-stdout"
	_ "go.jcbhmr.com/sys/wasip2/clocks/monotonic-clock"
	_ "go.jcbhmr.com/sys/wasip2/clocks/timezone"
	_ "go.jcbhmr.com/sys/wasip2/clocks/wall-clock"
	_ "go.jcbhmr.com/sys/wasip2/filesystem/preopens"
	_ "go.jcbhmr.com/sys/wasip2/filesystem/types"
	_ "go.jcbhmr.com/sys/wasip2/http/outgoing-handler"
	_ "go.jcbhmr.com/sys/wasip2/http/types"
	_ "go.jcbhmr.com/sys/wasip2/io/error"
	_ "go.jcbhmr.com/sys/wasip2/io/poll"
	_ "go.jcbhmr.com/sys/wasip2/io/streams"
	_ "go.jcbhmr.com/sys/wasip2/random/insecure"
	_ "go.jcbhmr.com/sys/wasip2/random/insecure-seed"
	_ "go.jcbhmr.com/sys/wasip2/random/random"
	_ "go.jcbhmr.com/sys/wasip2/sockets/instance-network"
	_ "go.jcbhmr.com/sys/wasip2/sockets/ip-name-lookup"
	_ "go.jcbhmr.com/sys/wasip2/sockets/network"
	_ "go.jcbhmr.com/sys/wasip2/sockets/tcp"
	_ "go.jcbhmr.com/sys/wasip2/sockets/tcp-create-socket"
	_ "go.jcbhmr.com/sys/wasip2/sockets/udp"
	_ "go.jcbhmr.com/sys/wasip2/sockets/udp-create-socket"
)

func TestImportAll(t *testing.T) {}
