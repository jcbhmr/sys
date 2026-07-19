//go:build wasip1 && wasm32

package wasip1

import (
	"structs"
	"unsafe"
)

type Size = uint32

type Filesize = uint64

type Timestamp = uint64

type Clockid uint32

const (
	ClockidRealtime Clockid = iota
	ClockidMonotonic
	ClockidProcessCputimeId
	ClockidThreadCputimeId
)

var clockidNames = [...]string{
	"realtime",
	"monotonic",
	"process_cputime_id",
	"thread_cputime_id",
}

func (e Clockid) String() string {
	return clockidNames[e]
}

type Errno uint16

const (
	ErrnoSuccess Errno = iota
	Errno2big
	ErrnoAcces
	ErrnoAddrinuse
	ErrnoAddrnotavail
	ErrnoAfnosupport
	ErrnoAgain
	ErrnoAlready
	ErrnoBadf
	ErrnoBadmsg
	ErrnoBusy
	ErrnoCanceled
	ErrnoChild
	ErrnoConnaborted
	ErrnoConnrefused
	ErrnoConnreset
	ErrnoDeadlk
	ErrnoDestaddrreq
	ErrnoDom
	ErrnoDquot
	ErrnoExist
	ErrnoFault
	ErrnoFbig
	ErrnoHostunreach
	ErrnoIdrm
	ErrnoIlseq
	ErrnoInprogress
	ErrnoIntr
	ErrnoInval
	ErrnoIo
	ErrnoIsconn
	ErrnoIsdir
	ErrnoLoop
	ErrnoMfile
	ErrnoMlink
	ErrnoMsgsize
	ErrnoMultihop
	ErrnoNametoolong
	ErrnoNetdown
	ErrnoNetreset
	ErrnoNetunreach
	ErrnoNfile
	ErrnoNobufs
	ErrnoNodev
	ErrnoNoent
	ErrnoNoexec
	ErrnoNolck
	ErrnoNolink
	ErrnoNomem
	ErrnoNomsg
	ErrnoNoprotoopt
	ErrnoNospc
	ErrnoNosys
	ErrnoNotconn
	ErrnoNotdir
	ErrnoNotempty
	ErrnoNotrecoverable
	ErrnoNotsock
	ErrnoNotsup
	ErrnoNotty
	ErrnoNxio
	ErrnoOverflow
	ErrnoOwnerdead
	ErrnoPerm
	ErrnoPipe
	ErrnoProto
	ErrnoProtonosupport
	ErrnoPrototype
	ErrnoRange
	ErrnoRofs
	ErrnoSpipe
	ErrnoSrch
	ErrnoStale
	ErrnoTimedout
	ErrnoTxtbsy
	ErrnoXdev
	ErrnoNotcapable
)

var errnoNames = [...]string{
	"success",
	"2big",
	"acces",
	"addrinuse",
	"addrnotavail",
	"afnosupport",
	"again",
	"already",
	"badf",
	"badmsg",
	"busy",
	"canceled",
	"child",
	"connaborted",
	"connrefused",
	"connreset",
	"deadlk",
	"destaddrreq",
	"dom",
	"dquot",
	"exist",
	"fault",
	"fbig",
	"hostunreach",
	"idrm",
	"ilseq",
	"inprogress",
	"intr",
	"inval",
	"io",
	"isconn",
	"isdir",
	"loop",
	"mfile",
	"mlink",
	"msgsize",
	"multihop",
	"nametoolong",
	"netdown",
	"netreset",
	"netunreach",
	"nfile",
	"nobufs",
	"nodev",
	"noent",
	"noexec",
	"nolck",
	"nolink",
	"nomem",
	"nomsg",
	"noprotoopt",
	"nospc",
	"nosys",
	"notconn",
	"notdir",
	"notempty",
	"notrecoverable",
	"notsock",
	"notsup",
	"notty",
	"nxio",
	"overflow",
	"ownerdead",
	"perm",
	"pipe",
	"proto",
	"protonosupport",
	"prototype",
	"range",
	"rofs",
	"spipe",
	"srch",
	"stale",
	"timedout",
	"txtbsy",
	"xdev",
	"notcapable",
}

func (e Errno) String() string {
	return errnoNames[e]
}

func (e Errno) Error() string {
	return e.String()
}

type Rights = uint64

const (
	RightsFdDatasync Rights = 1 << iota
	RightsFdRead
	RightsFdSeek
	RightsFdFdstatSetFlags
	RightsFdSync
	RightsFdTell
	RightsFdWrite
	RightsFdAdvise
	RightsFdAllocate
	RightsPathCreateFirectory
	RightsPathCreateFile
	RightsPathLinkSource
	RightsPathLinkTarget
	RightsPathOpen
	RightsFdReaddir
	RightsPathReadlink
	RightsPathRenameSource
	RightsPathRenameTarget
	RightsPathFilestatGet
	RightsPathFilestatSetSize
	RightsPathFilestatSetTimes
	RightsFdFilestatGet
	RightsFdFilestatSetSize
	RightsFdFilestatSetTimes
	RightsPathSymlink
	RightsPathRemoveDirectory
	RightsPathUnlinkFile
	RightsPollFdReadwrite
	RightsSockShutdown
	RightsSockAccept
)

type Fd = uint32

type Iovec struct {
	_ structs.HostLayout
	Buf    *uint8
	BufLen Size
}

type Ciovec struct {
	_ structs.HostLayout
	Buf    *uint8
	BufLen Size
}

type IovecArray = []Iovec

type CiovecArray = []Ciovec

type Filedelta = int64

type Whence uint8

const (
	WhenceSet Whence = iota
	WhenceCur
	WhenceEnd
)

var whenceNames = [...]string{
	"set",
	"cut",
	"end",
}

func (e Whence) String() string {
	return whenceNames[e]
}

type Dircookie = uint64

type Dirnamlen = uint32

type Inode = uint64

type Filetype uint8

const (
	FiletypeUnknown Filetype = iota
	FiletypeBlockDevice
	FiletypeCharacterDevice
	FiletypeDirectory
	FiletypeRegularFile
	FiletypeSocketDgram
	FiletypeSocketStream
	FiletypeSymbolicLink
)

var filetypeNames = [...]string{
	"unknown",
	"block_device",
	"character_device",
	"directory",
	"regular_file",
	"socket_dgram",
	"socket_stream",
	"symbolic_link",
}

func (e Filetype) String() string {
	return filetypeNames[e]
}

type Dirent struct {
	_ structs.HostLayout
	DNext   Dircookie
	DIno    Inode
	DNamlen Dirnamlen
	DType   Filetype
}

type Advice uint8

const (
	AdviceNormal Advice = iota
	AdviceSequential
	AdviceRandom
	AdviceWillneed
	AdviceDontneed
	AdviceNoreuse
)

var adviceNames = [...]string{
	"normal",
	"sequential",
	"random",
	"willneed",
	"dontneed",
	"noreuse",
}

func (e Advice) String() string {
	return adviceNames[e]
}

type Fdflags = uint16

const (
	FdflagsAppend Fdflags = 1 << iota
	FdflagsDsync
	FdflagsNonblock
	FdflagsRsync
	FdflagsSync
)

type Fdstat struct {
	_ structs.HostLayout
	FsFiletype         Filetype
	FsFlags            Fdflags
	FsRightsBase       Rights
	FsRightsInheriting Rights
}

type Device = uint64

type Fstflags = uint16

const (
	FstflagsAtim Fstflags = 1 << iota
	FstflagsAtimNow
	FstflagsMtim
	FstflagsMtimNow
)

type Lookupflags = uint32

const (
	LookupflagsSymlinkFollow Lookupflags = 1 << iota
)

type Oflags = uint16

const (
	OflagsCreat Oflags = 1 << iota
	OflagsDirectory
	OflagsExcl
	OflagsTrunc
)

type Linkcount = uint64

type Filestat struct {
	_ structs.HostLayout
	Dev      Device
	Ino      Inode
	Filetype Filetype
	Nlink    Linkcount
	Size     Filesize
	Atim     Timestamp
	Mtim     Timestamp
	Ctim     Timestamp
}

type Userdata = uint64

type Eventtype uint8

const (
	EventtypeClock Eventtype = iota
	EventtypeFdRead
	EventtypeFdWrite
)

var eventtypeNames = [...]string{
	"clock",
	"fd_read",
	"fd_write",
}

func (e Eventtype) String() string {
	return eventtypeNames[e]
}

type Eventrwflags = uint16

const (
	EventrwflagsFdReadwriteHangup Eventrwflags = 1 << iota
)

type EventFdReadwrite struct {
	_ structs.HostLayout
	Nbytes Filesize
	Flags  Eventrwflags
}

type Event struct {
	_ structs.HostLayout
	Userdata    Userdata
	Error       Errno
	Type        Eventtype
	FdReadwrite EventFdReadwrite
}

type Subclockflags = uint16

const (
	SubclockflagsSubscriptionClockAbstime Subclockflags = 1 << iota
)

type SubscriptionClock struct {
	_ structs.HostLayout
	Id        Clockid
	Timeout   Timestamp
	Precision Timestamp
	Flags     Subclockflags
}

type SubscriptionFdReadwrite struct {
	_ structs.HostLayout
	FileDescriptor Fd
}

type SubscriptionU struct {
	_ structs.HostLayout
	Tag  Eventtype
	Data [max(
		unsafe.Sizeof(SubscriptionClock{}),
		unsafe.Sizeof(SubscriptionFdReadwrite{}),
		unsafe.Sizeof(SubscriptionFdReadwrite{}),
	)]byte
}

func (u *SubscriptionU) Clock() *SubscriptionClock {
	if u.Tag != EventtypeClock {
		return nil
	}
	return (*SubscriptionClock)(unsafe.Pointer(&u.Data))
}

func (u *SubscriptionU) FdRead() *SubscriptionFdReadwrite {
	if u.Tag != EventtypeFdRead {
		return nil
	}
	return (*SubscriptionFdReadwrite)(unsafe.Pointer(&u.Data))
}

func (u *SubscriptionU) FdWrite() *SubscriptionFdReadwrite {
	if u.Tag != EventtypeFdWrite {
		return nil
	}
	return (*SubscriptionFdReadwrite)(unsafe.Pointer(&u.Data))
}

type Subscription struct {
	_ structs.HostLayout
	Userdata Userdata
	U        SubscriptionU
}

type Exitcode = uint32

type Signal uint8

const (
	SignalNone Signal = iota
	SignalHup
	SignalInt
	SignalQuit
	SignalIll
	SignalTrap
	SignalAbrt
	SignalBus
	SignalFpe
	SignalKill
	SignalUsr1
	SignalSegv
	SignalUsr2
	SignalPipe
	SignalAlrm
	SignalTerm
	SignalChld
	SignalCont
	SignalStop
	SignalTstp
	SignalTtin
	SignalTtou
	SignalUrg
	SignalXcpu
	SignalXfsz
	SignalVtalrm
	SignalProf
	SignalWinch
	SignalPoll
	SignalPwr
	SignalSys
)

var signalNames = [...]string{
    "none",
    "hup",
    "int",
    "quit",
    "ill",
    "trap",
    "abrt",
    "bus",
    "fpe",
    "kill",
    "usr1",
    "segv",
    "usr2",
    "pipe",
    "alrm",
    "term",
    "chld",
    "cont",
    "stop",
    "tstp",
    "ttin",
    "ttou",
    "urg",
    "xcpu",
    "xfsz",
    "vtalrm",
    "prof",
    "winch",
    "poll",
    "pwr",
    "sys",
}

func (e Signal) String() string {
	return signalNames[e]
}

type Riflags = uint16

const (
	RiflagsRecvPeek Riflags = 1 << iota
	RiflagsRecvWaitall
)

type Roflags = uint16

const (
	RoflagsRecvDataTruncated Roflags = 1 << iota
)

type Siflags = uint16

type Sdflags = uint8

const (
	SdflagsRd Sdflags = 1 << iota
	SdflagsWr
)

type Preopentype uint8

const (
	PreopentypeDir Preopentype = iota
)

var preopentypeNames = [...]string{
	"dir",
}

func (e Preopentype) String() string {
	return preopentypeNames[e]
}

type PrestatDir struct {
	_ structs.HostLayout
	PrNameLen Size
}

type Prestat struct {
	_ structs.HostLayout
	Tag  Preopentype
	Data [max(
		unsafe.Sizeof(PrestatDir{}),
	)]byte
}

func (u *Prestat) Dir() *PrestatDir {
	if u.Tag != PreopentypeDir {
		return nil
	}
	return (*PrestatDir)(unsafe.Pointer(&u.Data))
}
