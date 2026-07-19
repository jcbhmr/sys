//go:build wasip1 && wasm32

package wasip1

//go:wasmimport wasi_snapshot_preview1 args_get
//go:noescape
func wasmimport_args_get(argv **uint8, argvBuf *uint8) uint32

//go:wasmimport wasi_snapshot_preview1 args_sizes_get
//go:noescape
func wasmimport_args_sizes_get(rp0 *Size, rp1 *Size) uint32

//go:wasmimport wasi_snapshot_preview1 environ_get
//go:noescape
func wasmimport_environ_get(environ **uint8, environBuf *uint8) uint32

//go:wasmimport wasi_snapshot_preview1 environ_sizes_get
//go:noescape
func wasmimport_environ_sizes_get(rp0 *Size, rp1 *Size) uint32

//go:wasmimport wasi_snapshot_preview1 clock_res_get
//go:noescape
func wasmimport_clock_res_get(id Clockid, rp0 *Timestamp) uint32

//go:wasmimport wasi_snapshot_preview1 clock_time_get
//go:noescape
func wasmimport_clock_time_get(id Clockid, precision Timestamp, rp0 *Timestamp) uint32

//go:wasmimport wasi_snapshot_preview1 fd_advise
func wasmimport_fd_advise(fd Fd, offset Filesize, len Filesize, advice uint32) uint32

//go:wasmimport wasi_snapshot_preview1 fd_allocate
func wasmimport_fd_allocate(fd Fd, offset Filesize, len Filesize) uint32

//go:wasmimport wasi_snapshot_preview1 fd_close
func wasmimport_fd_close(fd Fd) uint32

//go:wasmimport wasi_snapshot_preview1 fd_datasync
func wasmimport_fd_datasync(fd Fd) uint32

//go:wasmimport wasi_snapshot_preview1 fd_fdstat_get
//go:noescape
func wasmimport_fd_fdstat_get(fd Fd, rp0 *Fdstat) uint32

//go:wasmimport wasi_snapshot_preview1 fd_fdstat_set_flags
func wasmimport_fd_fdstat_set_flags(fd Fd, flags uint32) uint32

//go:wasmimport wasi_snapshot_preview1 fd_fdstat_set_rights
func wasmimport_fd_fdstat_set_rights(fd Fd, fsRightsBase Rights, fsRightsInheriting Rights) uint32

//go:wasmimport wasi_snapshot_preview1 fd_filestat_get
//go:noescape
func wasmimport_fd_filestat_get(fd Fd, rp0 *Filestat) uint32

//go:wasmimport wasi_snapshot_preview1 fd_filestat_set_size
func wasmimport_fd_filestat_set_size(fd Fd, size Filesize) uint32

//go:wasmimport wasi_snapshot_preview1 fd_filestat_set_times
func wasmimport_fd_filestat_set_times(fd Fd, atim Timestamp, mtim Timestamp, fstFlags uint32) uint32

//go:wasmimport wasi_snapshot_preview1 fd_pread
//go:noescape
func wasmimport_fd_pread(fd Fd, iovsData *Iovec, iovsLen uint32, offset Filesize, rp0 *Size) uint32

//go:wasmimport wasi_snapshot_preview1 fd_prestat_get
//go:noescape
func wasmimport_fd_prestat_get(fd Fd, rp0 *Prestat) uint32

//go:wasmimport wasi_snapshot_preview1 fd_prestat_dir_name
//go:noescape
func wasmimport_fd_prestat_dir_name(fd Fd, path *uint8, pathLen Size) uint32

//go:wasmimport wasi_snapshot_preview1 fd_pwrite
//go:noescape
func wasmimport_fd_pwrite(fd Fd, iovsData *Ciovec, iovsLen uintptr, offset Filesize, rp0 *Size) uint32

//go:wasmimport wasi_snapshot_preview1 fd_read
//go:noescape
func wasmimport_fd_read(fd Fd, iovsData *Iovec, iovsLen uintptr, rp0 *Size) uint32

//go:wasmimport wasi_snapshot_preview1 fd_readdir
//go:noescape
func wasmimport_fd_readdir(fd Fd, buf *uint8, bufLen Size, cookie Dircookie, rp0 *Size) uint32

//go:wasmimport wasi_snapshot_preview1 fd_renumber
func wasmimport_fd_renumber(fd Fd, to Fd) uint32

//go:wasmimport wasi_snapshot_preview1 fd_seek
//go:noescape
func wasmimport_fd_seek(fd Fd, offset Filedelta, whence uint32, rp0 *Filesize) uint32

//go:wasmimport wasi_snapshot_preview1 fd_sync
func wasmimport_fd_sync(fd Fd) uint32

//go:wasmimport wasi_snapshot_preview1 fd_tell
//go:noescape
func wasmimport_fd_tell(fd Fd, rp0 *Filesize) uint32

//go:wasmimport wasi_snapshot_preview1 fd_write
//go:noescape
func wasmimport_fd_write(fd Fd, iovsData *Ciovec, iovsLen uintptr, rp0 *Size) uint32

//go:wasmimport wasi_snapshot_preview1 path_create_directory
//go:noescape
func wasmimport_path_create_directory(fd Fd, path string) uint32

//go:wasmimport wasi_snapshot_preview1 path_filestat_get
//go:noescape
func wasmimport_path_filestat_get(fd Fd, flags Lookupflags, path string, rp0 *Filestat) uint32

//go:wasmimport wasi_snapshot_preview1 path_filestat_set_times
//go:noescape
func wasmimport_path_filestat_set_times(fd Fd, flags Lookupflags, path string, atim Timestamp, mtim Timestamp, fstFlags uint32) uint32

//go:wasmimport wasi_snapshot_preview1 path_link
//go:noescape
func wasmimport_path_link(oldFd Fd, oldFlags Lookupflags, oldPath string, newFd Fd, newPath string) uint32

//go:wasmimport wasi_snapshot_preview1 path_open
//go:noescape
func wasmimport_path_open(fd Fd, dirflags Lookupflags, path string, oflags uint32, fsRights Rights, fsRightsInheriting Rights, fdflags uint32, rp0 *Fd) uint32

//go:wasmimport wasi_snapshot_preview1 path_readlink
//go:noescape
func wasmimport_path_readlink(fd Fd, path string, buf *uint8, bufLen Size, rp0 *Size) uint32

//go:wasmimport wasi_snapshot_preview1 path_remove_directory
//go:noescape
func wasmimport_path_remove_directory(fd Fd, path string) uint32

//go:wasmimport wasi_snapshot_preview1 path_rename
//go:noescape
func wasmimport_path_rename(fd Fd, oldPath string, newFd Fd, newPath string) uint32

//go:wasmimport wasi_snapshot_preview1 path_symlink
//go:noescape
func wasmimport_path_symlink(oldPath string, fd Fd, newPath string) uint32

//go:wasmimport wasi_snapshot_preview1 path_unlink_file
//go:noescape
func wasmimport_path_unlink_file(fd Fd, path string) uint32

//go:wasmimport wasi_snapshot_preview1 poll_oneoff
//go:noescape
func wasmimport_poll_oneoff(in *Subscription, out *Event, nsubscriptions Size, rp0 *Size) uint32

//go:wasmimport wasi_snapshot_preview1 proc_exit
func wasmimport_proc_exit(rval Exitcode)

//go:wasmimport wasi_snapshot_preview1 proc_raise
func wasmimport_proc_raise(sig uint32) uint32

//go:wasmimport wasi_snapshot_preview1 sched_yield
func wasmimport_sched_yield() uint32

//go:wasmimport wasi_snapshot_preview1 random_get
//go:noescape
func wasmimport_random_get(buf *uint8, bufLen Size) uint32

//go:wasmimport wasi_snapshot_preview1 sock_accept
//go:noescape
func wasmimport_sock_accept(fd Fd, flags uint32, rp0 *Fd) uint32

//go:wasmimport wasi_snapshot_preview1 sock_recv
//go:noescape
func wasmimport_sock_recv(fd Fd, riDataData *Iovec, riDataLen uintptr, riFlags uint32, rp0 *Size, rp1 *Roflags) uint32

//go:wasmimport wasi_snapshot_preview1 sock_send
//go:noescape
func wasmimport_sock_send(fd Fd, siDataData *Ciovec, siDataLen uintptr, siFlags uint32, rp0 *Size) uint32

//go:wasmimport wasi_snapshot_preview1 sock_shutdown
func wasmimport_sock_shutdown(fd Fd, how uint32) uint32
