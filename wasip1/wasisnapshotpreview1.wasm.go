//go:build wasip1

package wasip1

//go:wasmimport wasi_snapshot_preview1 args_get
//go:noescape
func wasmimport_args_get(argv **uint8, argvBuf *uint8) Errno

//go:wasmimport wasi_snapshot_preview1 args_sizes_get
//go:noescape
func wasmimport_args_sizes_get(rp0 *Size, rp1 *Size) Errno

//go:wasmimport wasi_snapshot_preview1 environ_get
//go:noescape
func wasmimport_environ_get(environ **uint8, environBuf *uint8) Errno

//go:wasmimport wasi_snapshot_preview1 environ_sizes_get
//go:noescape
func wasmimport_environ_sizes_get(rp0 *Size, rp1 *Size) Errno

//go:wasmimport wasi_snapshot_preview1 clock_res_get
//go:noescape
func wasmimport_clock_res_get(id Clockid, rp0 *Timestamp) Errno

//go:wasmimport wasi_snapshot_preview1 clock_time_get
//go:noescape
func wasmimport_clock_time_get(id Clockid, precision Timestamp, rp0 *Timestamp) Errno

//go:wasmimport wasi_snapshot_preview1 fd_advise
func wasmimport_fd_advise(fd Fd, offset Filesize, len Filesize, advice Advice) Errno
