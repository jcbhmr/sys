//go:build wasip1 && wasm32

package wasip1

import "unsafe"

func ArgsGet(argv **uint8, argvBuf *uint8) Errno {
	return Errno(wasmimport_args_get(argv, argvBuf))
}

func ArgsSizesGet() (r0 Size, r1 Size, errno Errno) {
	errno = Errno(wasmimport_args_sizes_get(&r0, &r1))
	return r0, r1, errno
}

func EnvironGet(environ **uint8, environBuf *uint8) Errno {
	return Errno(wasmimport_environ_get(environ, environBuf))
}

func EnvironSizesGet() (r0 Size, r1 Size, errno Errno) {
	errno = Errno(wasmimport_environ_sizes_get(&r0, &r1))
	return r0, r1, errno
}

func ClockResGet(id Clockid) (r0 Timestamp, errno Errno) {
	errno = Errno(wasmimport_clock_res_get(id, &r0))
	return r0, errno
}

func ClockTimeGet(id Clockid, precision Timestamp) (r0 Timestamp, errno Errno) {
	errno = Errno(wasmimport_clock_time_get(id, precision, &r0))
	return r0, errno
}

func FdAdvise(fd Fd, offset Filesize, len Filesize, advice Advice) Errno {
	return Errno(wasmimport_fd_advise(fd, offset, len, uint32(advice)))
}

func FdAllocate(fd Fd, offset Filesize, len Filesize) Errno {
	return Errno(wasmimport_fd_allocate(fd, offset, len))
}

func FdClose(fd Fd) Errno {
	return Errno(wasmimport_fd_close(fd))
}

func FdDatasync(fd Fd) Errno {
	return Errno(wasmimport_fd_datasync(fd))
}

func FdFdstatGet(fd Fd) (r0 Fdstat, errno Errno) {
	errno = Errno(wasmimport_fd_fdstat_get(fd, &r0))
	return r0, errno
}

func FdFdstatSetFlags(fd Fd, flags Fdflags) Errno {
	return Errno(wasmimport_fd_fdstat_set_flags(fd, uint32(flags)))
}

func FdFdstatSetRights(fd Fd, fsRightsBase Rights, fsRightsInheriting Rights) Errno {
	return Errno(wasmimport_fd_fdstat_set_rights(fd, fsRightsBase, fsRightsInheriting))
}

func FdFilestatGet(fd Fd) (r0 Filestat, errno Errno) {
	errno = Errno(wasmimport_fd_filestat_get(fd, &r0))
	return r0, errno
}

func FdFilestatSetSize(fd Fd, size Filesize) Errno {
	return Errno(wasmimport_fd_filestat_set_size(fd, size))
}

func FdFilestatSetTimes(fd Fd, atim Timestamp, mtim Timestamp, fstFlags Fstflags) Errno {
	return Errno(wasmimport_fd_filestat_set_times(fd, atim, mtim, uint32(fstFlags)))
}

func FdPread(fd Fd, iovs IovecArray, offset Filesize) (r0 Size, errno Errno) {
	errno = Errno(wasmimport_fd_pread(fd, unsafe.SliceData(iovs), uint32(len(iovs)), offset, &r0))
	return r0, errno
}

func FdPrestatGet(fd Fd) (r0 Prestat, errno Errno) {
	errno = Errno(wasmimport_fd_prestat_get(fd, &r0))
	return r0, errno
}

func FdPrestatDirName(fd Fd, path *uint8, pathLen Size) Errno {
	return Errno(wasmimport_fd_prestat_dir_name(fd, path, pathLen))
}

func FdPwrite(fd Fd, iovs CiovecArray, offset Filesize) (r0 Size, errno Errno) {
	errno = Errno(wasmimport_fd_pwrite(fd, unsafe.SliceData(iovs), uintptr(len(iovs)), offset, &r0))
	return r0, errno
}

func FdRead(fd Fd, iovs IovecArray) (r0 Size, errno Errno) {
	errno = Errno(wasmimport_fd_read(fd, unsafe.SliceData(iovs), uintptr(len(iovs)), &r0))
	return r0, errno
}

func FdReaddir(fd Fd, buf *uint8, bufLen Size, cookie Dircookie) (r0 Size, errno Errno) {
	errno = Errno(wasmimport_fd_readdir(fd, buf, bufLen, cookie, &r0))
	return r0, errno
}

func FdRenumber(fd Fd, to Fd) Errno {
	return Errno(wasmimport_fd_renumber(fd, to))
}

func FdSeek(fd Fd, offset Filedelta, whence Whence) (r0 Filesize, errno Errno) {
	errno = Errno(wasmimport_fd_seek(fd, offset, uint32(whence), &r0))
	return r0, errno
}

func FdSync(fd Fd) Errno {
	return Errno(wasmimport_fd_sync(fd))
}

func FdTell(fd Fd) (r0 Filesize, errno Errno) {
	errno = Errno(wasmimport_fd_tell(fd, &r0))
	return r0, errno
}

func FdWrite(fd Fd, iovs CiovecArray) (r0 Size, errno Errno) {
	errno = Errno(wasmimport_fd_write(fd, unsafe.SliceData(iovs), uintptr(len(iovs)), &r0))
	return r0, errno
}

func PathCreateDirectory(fd Fd, path string) Errno {
	return Errno(wasmimport_path_create_directory(fd, path))
}

func PathFilestatGet(fd Fd, flags Lookupflags, path string) (r0 Filestat, errno Errno) {
	errno = Errno(wasmimport_path_filestat_get(fd, flags, path, &r0))
	return r0, errno
}

func PathFilestatSetTimes(fd Fd, flags Lookupflags, path string, atim Timestamp, mtim Timestamp, fstFlags Fstflags) Errno {
	return Errno(wasmimport_path_filestat_set_times(fd, flags, path, atim, mtim, uint32(fstFlags)))
}

func PathLink(oldFd Fd, oldFlags Lookupflags, oldPath string, newFd Fd, newPath string) Errno {
	return Errno(wasmimport_path_link(oldFd, oldFlags, oldPath, newFd, newPath))
}

func PathOpen(fd Fd, dirflags Lookupflags, path string, oflags Oflags, fsRightsBase Rights, fsRightsInheriting Rights, fdflags Fdflags) (r0 Fd, errno Errno) {
	errno = Errno(wasmimport_path_open(fd, dirflags, path, uint32(oflags), fsRightsBase, fsRightsInheriting, uint32(fdflags), &r0))
	return r0, errno
}

func PathReadlink(fd Fd, path string, buf *uint8, bufLen Size) (r0 Size, errno Errno) {
	errno = Errno(wasmimport_path_readlink(fd, path, buf, bufLen, &r0))
	return r0, errno
}

func PathRemoveDirectory(fd Fd, path string) Errno {
	return Errno(wasmimport_path_remove_directory(fd, path))
}

func PathRename(fd Fd, oldPath string, newFd Fd, newPath string) Errno {
	return Errno(wasmimport_path_rename(fd, oldPath, newFd, newPath))
}

func PathSymlink(oldPath string, fd Fd, newPath string) Errno {
	return Errno(wasmimport_path_symlink(oldPath, fd, newPath))
}

func PathUnlinkFile(fd Fd, path string) Errno {
	return Errno(wasmimport_path_unlink_file(fd, path))
}

func PollOneoff(in *Subscription, out *Event, nsubscriptions Size) (r0 Size, errno Errno) {
	errno = Errno(wasmimport_poll_oneoff(in, out, nsubscriptions, &r0))
	return r0, errno
}

func ProcExit(rval Exitcode) {
	wasmimport_proc_exit(rval)
	panic("noreturn")
}

func ProcRaise(sig Signal) Errno {
	return Errno(wasmimport_proc_raise(uint32(sig)))
}

func SchedYield() Errno {
	return Errno(wasmimport_sched_yield())
}

func RandomGet(buf *uint8, bufLen Size) Errno {
	return Errno(wasmimport_random_get(buf, bufLen))
}

func SockAccept(fd Fd, flags Fdflags) (r0 Fd, errno Errno) {
	errno = Errno(wasmimport_sock_accept(fd, uint32(flags), &r0))
	return r0, errno
}

func SockRecv(fd Fd, riData IovecArray, riFlags Riflags) (r0 Size, r1 Roflags, errno Errno) {
	errno = Errno(wasmimport_sock_recv(fd, unsafe.SliceData(riData), uintptr(len(riData)), uint32(riFlags), &r0, &r1))
	return r0, r1, errno
}

func SockSend(fd Fd, siData CiovecArray, siFlags Siflags) (r0 Size, errno Errno) {
	errno = Errno(wasmimport_sock_send(fd, unsafe.SliceData(siData), uintptr(len(siData)), uint32(siFlags), &r0))
	return r0, errno
}

func SockShutdown(fd Fd, how Sdflags) Errno {
	return Errno(wasmimport_sock_shutdown(fd, uint32(how)))
}
