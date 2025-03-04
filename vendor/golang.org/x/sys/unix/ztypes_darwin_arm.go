// NOTE: cgo can't generate struct Stat_t and struct Statfs_t yet
// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs types_darwin.go

// +build arm,darwin

package unix

const (
	SizeofPtr      = 0x4
	SizeofShort    = 0x2
	SizeofInt      = 0x4
	SizeofLong     = 0x4
	SizeofLongLong = 0x8
)

type (
	_C_short int16
	_C_int int32
	_C_long int32
	_C_long_long int64
)

type Timespec struct {
	Sec  int32
	Nsec int32
}

type Timeval struct {
	Sec  int32
	Usec int32
}

type Timeval32 [0]byte

type Rusage struct {
	Utime    Timeval
	Stime    Timeval
	Maxrss   int32
	Ixrss    int32
	Idrss    int32
	Isrss    int32
	Minflt   int32
	Majflt   int32
	Nswap    int32
	Inblock  int32
	Oublock  int32
	Msgsnd   int32
	Msgrcv   int32
	Nsignals int32
	Nvcsw    int32
	Nivcsw   int32
}

type Rlimit struct {
	Cur uint64
	Max uint64
}

type _Gid_t uint32

type Stat_t struct {
	Dev     int32
	Mode    uint16
	Nlink   uint16
	Ino     uint64
	Uid     uint32
	Gid     uint32
	Rdev    int32
	Atim    Timespec
	Mtim    Timespec
	Ctim    Timespec
	Btim    Timespec
	Size    int64
	Blocks  int64
	Blksize int32
	Flags   uint32
	Gen     uint32
	Lspare  int32
	Qspare  [2]int64
}

type Statfs_t struct {
	Bsize       uint32
	Iosize      int32
	Blocks      uint64
	Bfree       uint64
	Bavail      uint64
	Files       uint64
	Ffree       uint64
	Fsid        Fsid
	Owner       uint32
	Type        uint32
	Flags       uint32
	Fssubtype   uint32
	Fstypename  [16]int8
	Mntonname   [1024]int8
	Mntfromname [1024]int8
	Reserved    [8]uint32
}

type Flock_t struct {
	Start  int64
	Len    int64
	Pid    int32
	Type   int16
	Whence int16
}

type Fstore_t struct {
	Flags      uint32
	Posmode    int32
	Offset     int64
	Length     int64
	Bytesalloc int64
}

type Radvisory_t struct {
	Offset int64
	Count  int32
}

type Fbootstraptransfer_t struct {
	Offset int64
	Length uint32
	Buffer *byte
}

type Log2phys_t struct {
	Flags       uint32
	Contigbytes int64
	Devoffset   int64
}

type Fsid struct {
	Val [2]int32
}

type Dirent struct {
	Ino     uint64
	Seekoff uint64
	Reclen  uint16
	Namlen  uint16
	Type    uint8
	Name    [1024]int8
	_       [3]byte
}

type RawSockaddrInet4 struct {
	Len    uint8
	Family uint8
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]int8
}

type RawSockaddrInet6 struct {
	Len      uint8
	Family   uint8
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte /* in6_addr */
	Scope_id uint32
}

type RawSockaddrUnix struct {
	Len    uint8
	Family uint8
	Path   [104]int8
}

type RawSockaddrDatalink struct {
	Len    uint8
	Family uint8
	Index  uint16
	Type   uint8
	Nlen   uint8
	Alen   uint8
	Slen   uint8
	Data   [12]int8
}

type RawSockaddr struct {
	Len    uint8
	Family uint8
	Data   [14]int8
}

type RawSockaddrAny struct {
	Addr RawSockaddr
	Pad  [92]int8
}

type _Socklen uint32

type Linger struct {
	Onoff  int32
	Linger int32
}

type Iovec struct {
	Base *byte
	Len  uint32
}

type IPMreq struct {
	Multiaddr [4]byte /* in_addr */
	Interface [4]byte /* in_addr */
}

type IPv6Mreq struct {
	Multiaddr [16]byte /* in6_addr */
	Interface uint32
}

type Msghdr struct {
	Name       *byte
	Namelen    uint32
	Iov        *Iovec
	Iovlen     int32
	Control    *byte
	Controllen uint32
	Flags      int32
}

type Cmsghdr struct {
	Len   uint32
	Level int32
	Type  int32
}

type Inet4Pktinfo struct {
	Ifindex  uint32
	Spec_dst [4]byte /* in_addr */
	Addr     [4]byte /* in_addr */
}

type Inet6Pktinfo struct {
	Addr    [16]byte /* in6_addr */
	Ifindex uint32
}

type IPv6MTUInfo struct {
	Addr RawSockaddrInet6
	Mtu  uint32
}

type ICMPv6Filter struct {
	Filt [8]uint32
}

const (
	SizeofSockaddrInet4    = 0x10
	SizeofSockaddrInet6    = 0x1c
	SizeofSockaddrAny      = 0x6c
	SizeofSockaddrUnix     = 0x6a
	SizeofSockaddrDatalink = 0x14
	SizeofLinger           = 0x8
	SizeofIPMreq           = 0x8
	SizeofIPv6Mreq         = 0x14
	SizeofMsghdr           = 0x1c
	SizeofCmsghdr          = 0xc
	SizeofInet4Pktinfo     = 0xc
	SizeofInet6Pktinfo     = 0x14
	SizeofIPv6MTUInfo      = 0x20
	SizeofICMPv6Filter     = 0x20
)

const (
	PTRACE_TRACEME = 0x0
	PTRACE_CONT    = 0x7
	PTRACE_KILL    = 0x8
)

type Kevent_t struct {
	Ident  uint32
	Filter int16
	Flags  uint16
	Fflags uint32
	Data   int32
	Udata  *byte
}

type FdSet struct {
	Bits [32]int32
}

const (
	SizeofIfMsghdr    = 0x70
	SizeofIfData      = 0x60
	SizeofIfaMsghdr   = 0x14
	SizeofIfmaMsghdr  = 0x10
	SizeofIfmaMsghdr2 = 0x14
	SizeofRtMsghdr    = 0x5c
	SizeofRtMetrics   = 0x38
)

type IfMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Addrs   int32
	Flags   int32
	Index   uint16
	_       [2]byte
	Data    IfData
}

type IfData struct {
	Type       uint8
	Typelen    uint8
	Physical   uint8
	Addrlen    uint8
	Hdrlen     uint8
	Recvquota  uint8
	Xmitquota  uint8
	Unused1    uint8
	Mtu        uint32
	Metric     uint32
	Baudrate   uint32
	Ipackets   uint32
	Ierrors    uint32
	Opackets   uint32
	Oerrors    uint32
	Collisions uint32
	Ibytes     uint32
	Obytes     uint32
	Imcasts    uint32
	Omcasts    uint32
	Iqdrops    uint32
	Noproto    uint32
	Recvtiming uint32
	Xmittiming uint32
	Lastchange Timeval
	Unused2    uint32
	Hwassist   uint32
	Reserved1  uint32
	Reserved2  uint32
}

type IfaMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Addrs   int32
	Flags   int32
	Index   uint16
	_       [2]byte
	Metric  int32
}

type IfmaMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Addrs   int32
	Flags   int32
	Index   uint16
	_       [2]byte
}

type IfmaMsghdr2 struct {
	Msglen   uint16
	Version  uint8
	Type     uint8
	Addrs    int32
	Flags    int32
	Index    uint16
	_        [2]byte
	Refcount int32
}

type RtMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Index   uint16
	_       [2]byte
	Flags   int32
	Addrs   int32
	Pid     int32
	Seq     int32
	Errno   int32
	Use     int32
	Inits   uint32
	Rmx     RtMetrics
}

type RtMetrics struct {
	Locks    uint32
	Mtu      uint32
	Hopcount uint32
	Expire   int32
	Recvpipe uint32
	Sendpipe uint32
	Ssthresh uint32
	Rtt      uint32
	Rttvar   uint32
	Pksent   uint32
	Filler   [4]uint32
}

const (
	SizeofBpfVersion = 0x4
	SizeofBpfStat    = 0x8
	SizeofBpfProgram = 0x8
	SizeofBpfInsn    = 0x8
	SizeofBpfHdr     = 0x14
)

type BpfVersion struct {
	Major uint16
	Minor uint16
}

type BpfStat struct {
	Recv uint32
	Drop uint32
}

type BpfProgram struct {
	Len   uint32
	Insns *BpfInsn
}

type BpfInsn struct {
	Code uint16
	Jt   uint8
	Jf   uint8
	K    uint32
}

type BpfHdr struct {
	Tstamp  Timeval
	Caplen  uint32
	Datalen uint32
	Hdrlen  uint16
	_       [2]byte
}

type Termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Cc     [20]uint8
	Ispeed uint32
	Ospeed uint32
}

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

const (
	AT_FDCWD            = -0x2
	AT_REMOVEDIR        = 0x80
	AT_SYMLINK_FOLLOW   = 0x40
	AT_SYMLINK_NOFOLLOW = 0x20
)

type PollFd struct {
	Fd      int32
	Events  int16
	Revents int16
}

const (
	POLLERR    = 0x8
	POLLHUP    = 0x10
	POLLIN     = 0x1
	POLLNVAL   = 0x20
	POLLOUT    = 0x4
	POLLPRI    = 0x2
	POLLRDBAND = 0x80
	POLLRDNORM = 0x40
	POLLWRBAND = 0x100
	POLLWRNORM = 0x4
)

type Utsname struct {
	Sysname  [256]byte
	Nodename [256]byte
	Release  [256]byte
	Version  [256]byte
	Machine  [256]byte
}

const SizeofClockinfo = 0x14

type Clockinfo struct {
	Hz      int32
	Tick    int32
	Tickadj int32
	Stathz  int32
	Profhz  int32
}
