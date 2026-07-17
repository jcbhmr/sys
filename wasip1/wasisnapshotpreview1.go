//go:build wasip1

package wasip1

// Read command-line argument data.
// 
// The size of the array should match that returned by `args_sizes_get`.
// 
// Each argument is expected to be `\0` terminated.
// 
// The first argument should be a string containing the "name" of the
// program. This need not be a usable filesystem path or even file name,
// and may even be a fixed string. Subsequent arguments are the arguments
// passed to the program by the user.
func ArgsGet(argv **uint8, argvBuf *uint8) (Errno) {
}

// Return command-line argument data sizes.
func ArgsSizesGet() (RECORD, Errno) {
}

// Read environment variable data.
// The sizes of the buffers should match that returned by `environ_sizes_get`.
// Key/value pairs are expected to be joined with `=`s, and terminated with `\0`s.
func EnvironGet(environ **uint8, environBuf *uint8) (Errno) {
}

// Return environment variable data sizes.
func EnvironSizesGet() (RECORD, Errno) {
}

// Return the resolution of a clock.
// Implementations are required to provide a non-zero value for supported clocks. For unsupported clocks,
// return `errno::inval`.
// Note: This is similar to `clock_getres` in POSIX.
func ClockResGet(id Clockid) (Timestamp, Errno) {
}

// Return the time value of a clock.
// Note: This is similar to `clock_gettime` in POSIX.
func ClockTimeGet(id Clockid, precision Timestamp) (Timestamp, Errno) {
}

// Provide file advisory information on a file descriptor.
// Note: This is similar to `posix_fadvise` in POSIX.
func FdAdvise(fd Fd, offset Filesize, len Filesize, advice Advice) (Errno) {
}

// Force the allocation of space in a file.
// Note: This is similar to `posix_fallocate` in POSIX.
func FdAllocate(fd Fd, offset Filesize, len Filesize) (Errno) {
}

// Close a file descriptor.
// Note: This is similar to `close` in POSIX.
func FdClose(fd Fd) (Errno) {
}

// Synchronize the data of a file to disk.
// Note: This is similar to `fdatasync` in POSIX.
func FdDatasync(fd Fd) (Errno) {
}

// Get the attributes of a file descriptor.
// Note: This returns similar flags to `fcntl(fd, F_GETFL)` in POSIX, as well as additional fields.
func FdFdstatGet(fd Fd) (Fdstat, Errno) {
}

// Adjust the flags associated with a file descriptor.
// Note: This is similar to `fcntl(fd, F_SETFL, flags)` in POSIX.
func FdFdstatSetFlags(fd Fd, flags Fdflags) (Errno) {
}

// Adjust the rights associated with a file descriptor.
// This can only be used to remove rights, and returns `errno::notcapable` if called in a way that would attempt to add rights
func FdFdstatSetRights(fd Fd, fsRightsBase Rights, fsRightsInheriting Rights) (Errno) {
}

// Return the attributes of an open file.
func FdFilestatGet(fd Fd) (Filestat, Errno) {
}

// Adjust the size of an open file. If this increases the file's size, the extra bytes are filled with zeros.
// Note: This is similar to `ftruncate` in POSIX.
func FdFilestatSetSize(fd Fd, size Filesize) (Errno) {
}

// Adjust the timestamps of an open file or directory.
// Note: This is similar to `futimens` in POSIX.
func FdFilestatSetTimes(fd Fd, atim Timestamp, mtim Timestamp, fstFlags Fstflags) (Errno) {
}

// Read from a file descriptor, without using and updating the file descriptor's offset.
// Note: This is similar to `preadv` in Linux (and other Unix-es).
func FdPread(fd Fd, iovs IovecArray, offset Filesize) (Size, Errno) {
}

// Return a description of the given preopened file descriptor.
func FdPrestatGet(fd Fd) (Prestat, Errno) {
}

// Return a description of the given preopened file descriptor.
func FdPrestatDirName(fd Fd, path *uint8, pathLen Size) (Errno) {
}

// Write to a file descriptor, without using and updating the file descriptor's offset.
// Note: This is similar to `pwritev` in Linux (and other Unix-es).
// 
// Like Linux (and other Unix-es), any calls of `pwrite` (and other
// functions to read or write) for a regular file by other threads in the
// WASI process should not be interleaved while `pwrite` is executed.
func FdPwrite(fd Fd, iovs CiovecArray, offset Filesize) (Size, Errno) {
}

// Read from a file descriptor.
// Note: This is similar to `readv` in POSIX.
func FdRead(fd Fd, iovs IovecArray) (Size, Errno) {
}

// Read directory entries from a directory.
// When successful, the contents of the output buffer consist of a sequence of
// directory entries. Each directory entry consists of a `dirent` object,
// followed by `dirent::d_namlen` bytes holding the name of the directory
// entry.
// This function fills the output buffer as much as possible, potentially
// truncating the last directory entry. This allows the caller to grow its
// read buffer size in case it's too small to fit a single large directory
// entry, or skip the oversized directory entry.
// 
// Entries for the special `.` and `..` directory entries are included in the
// sequence.
func FdReaddir(fd Fd, buf *uint8, bufLen Size, cookie Dircookie) (Size, Errno) {
}

// Atomically replace a file descriptor by renumbering another file descriptor.
// Due to the strong focus on thread safety, this environment does not provide
// a mechanism to duplicate or renumber a file descriptor to an arbitrary
// number, like `dup2()`. This would be prone to race conditions, as an actual
// file descriptor with the same number could be allocated by a different
// thread at the same time.
// This function provides a way to atomically renumber file descriptors, which
// would disappear if `dup2()` were to be removed entirely.
func FdRenumber(fd Fd, to Fd) (Errno) {
}

// Move the offset of a file descriptor.
// Note: This is similar to `lseek` in POSIX.
func FdSeek(fd Fd, offset Filedelta, whence Whence) (Filesize, Errno) {
}

// Synchronize the data and metadata of a file to disk.
// Note: This is similar to `fsync` in POSIX.
func FdSync(fd Fd) (Errno) {
}

// Return the current offset of a file descriptor.
// Note: This is similar to `lseek(fd, 0, SEEK_CUR)` in POSIX.
func FdTell(fd Fd) (Filesize, Errno) {
}

// Write to a file descriptor.
// Note: This is similar to `writev` in POSIX.
// 
// Like POSIX, any calls of `write` (and other functions to read or write)
// for a regular file by other threads in the WASI process should not be
// interleaved while `write` is executed.
func FdWrite(fd Fd, iovs CiovecArray) (Size, Errno) {
}

// Create a directory.
// Note: This is similar to `mkdirat` in POSIX.
func PathCreateDirectory(fd Fd, path string) (Errno) {
}

// Return the attributes of a file or directory.
// Note: This is similar to `stat` in POSIX.
func PathFilestatGet(fd Fd, flags Lookupflags, path string) (Filestat, Errno) {
}

// Adjust the timestamps of a file or directory.
// Note: This is similar to `utimensat` in POSIX.
func PathFilestatSetTimes(fd Fd, flags Lookupflags, path string, atim Timestamp, mtim Timestamp, fstFlags Fstflags) (Errno) {
}

// Create a hard link.
// Note: This is similar to `linkat` in POSIX.
func PathLink(oldFd Fd, oldFlags Lookupflags, oldPath string, newFd Fd, newPath string) (Errno) {
}

// Open a file or directory.
// The returned file descriptor is not guaranteed to be the lowest-numbered
// file descriptor not currently open; it is randomized to prevent
// applications from depending on making assumptions about indexes, since this
// is error-prone in multi-threaded contexts. The returned file descriptor is
// guaranteed to be less than 2**31.
// Note: This is similar to `openat` in POSIX.
func PathOpen(fd Fd, dirflags Lookupflags, path string, oflags Oflags, fsRightsBase Rights, fsRightsInheriting Rights, fdflags Fdflags) (Fd, Errno) {
}

// Read the contents of a symbolic link.
// Note: This is similar to `readlinkat` in POSIX. If `buf` is not large
// enough to contain the contents of the link, the first `buf_len` bytes will be
// be written to `buf`.
func PathReadlink(fd Fd, path string, buf *uint8, bufLen Size) (Size, Errno) {
}

// Remove a directory.
// Return `errno::notempty` if the directory is not empty.
// Note: This is similar to `unlinkat(fd, path, AT_REMOVEDIR)` in POSIX.
func PathRemoveDirectory(fd Fd, path string) (Errno) {
}

// Rename a file or directory.
// Note: This is similar to `renameat` in POSIX.
func PathRename(fd Fd, oldPath string, newFd Fd, newPath string) (Errno) {
}

// Create a symbolic link.
// Note: This is similar to `symlinkat` in POSIX.
func PathSymlink(oldPath string, fd Fd, newPath string) (Errno) {
}

// Unlink a file.
// Return `errno::isdir` if the path refers to a directory.
// Note: This is similar to `unlinkat(fd, path, 0)` in POSIX.
func PathUnlinkFile(fd Fd, path string) (Errno) {
}

// Concurrently poll for the occurrence of a set of events.
// 
// If `nsubscriptions` is 0, returns `errno::inval`.
func PollOneoff(in *Subscription, out *Event, nsubscriptions Size) (Size, Errno) {
}

// Terminate the process normally. An exit code of 0 indicates successful
// termination of the program. The meanings of other values is dependent on
// the environment.
func ProcExit(rval Exitcode) {
}

// Send a signal to the process of the calling thread.
// Note: This is similar to `raise` in POSIX.
func ProcRaise(sig Signal) (Errno) {
}

// Temporarily yield execution of the calling thread.
// Note: This is similar to `sched_yield` in POSIX.
func SchedYield() (Errno) {
}

// Write high-quality random data into a buffer.
// This function blocks when the implementation is unable to immediately
// provide sufficient high-quality random data.
func RandomGet(buf *uint8, bufLen Size) (Errno) {
}

// Accept a new incoming connection.
// Note: This is similar to `accept` in POSIX.
func SockAccept(fd Fd, flags Fdflags) (Fd, Errno) {
}

// Receive a message from a socket.
// Note: This is similar to `recv` in POSIX, though it also supports reading
// the data into multiple buffers in the manner of `readv`.
func SockRecv(fd Fd, riData IovecArray, riFlags Riflags) (RECORD, Errno) {
}

// Send a message on a socket.
// Note: This is similar to `send` in POSIX, though it also supports writing
// the data from multiple buffers in the manner of `writev`.
func SockSend(fd Fd, siData CiovecArray, siFlags Siflags) (Size, Errno) {
}

// Shut down socket send and receive channels.
// Note: This is similar to `shutdown` in POSIX.
func SockShutdown(fd Fd, how Sdflags) (Errno) {
}

