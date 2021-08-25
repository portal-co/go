// mksyscall_libc.pl -solaris -tags solaris,amd64 syscall_solaris.go syscall_solaris_amd64.go
// Code generated by the command above; DO NOT EDIT.

//go:build solaris && amd64

package syscall

import "unsafe"

//go:cgo_import_dynamic libc_Getcwd getcwd "libc.so"
//go:cgo_import_dynamic libc_getgroups getgroups "libc.so"
//go:cgo_import_dynamic libc_setgroups setgroups "libc.so"
//go:cgo_import_dynamic libc_fcntl fcntl "libc.so"
//go:cgo_import_dynamic libc_accept accept "libsocket.so"
//go:cgo_import_dynamic libc___xnet_sendmsg __xnet_sendmsg "libsocket.so"
//go:cgo_import_dynamic libc_Access access "libc.so"
//go:cgo_import_dynamic libc_Adjtime adjtime "libc.so"
//go:cgo_import_dynamic libc_Chdir chdir "libc.so"
//go:cgo_import_dynamic libc_Chmod chmod "libc.so"
//go:cgo_import_dynamic libc_Chown chown "libc.so"
//go:cgo_import_dynamic libc_Chroot chroot "libc.so"
//go:cgo_import_dynamic libc_Close close "libc.so"
//go:cgo_import_dynamic libc_Dup dup "libc.so"
//go:cgo_import_dynamic libc_Fchdir fchdir "libc.so"
//go:cgo_import_dynamic libc_Fchmod fchmod "libc.so"
//go:cgo_import_dynamic libc_Fchown fchown "libc.so"
//go:cgo_import_dynamic libc_Fpathconf fpathconf "libc.so"
//go:cgo_import_dynamic libc_Fstat fstat "libc.so"
//go:cgo_import_dynamic libc_Getdents getdents "libc.so"
//go:cgo_import_dynamic libc_Getgid getgid "libc.so"
//go:cgo_import_dynamic libc_Getpid getpid "libc.so"
//go:cgo_import_dynamic libc_Geteuid geteuid "libc.so"
//go:cgo_import_dynamic libc_Getegid getegid "libc.so"
//go:cgo_import_dynamic libc_Getppid getppid "libc.so"
//go:cgo_import_dynamic libc_Getpriority getpriority "libc.so"
//go:cgo_import_dynamic libc_Getrlimit getrlimit "libc.so"
//go:cgo_import_dynamic libc_Gettimeofday gettimeofday "libc.so"
//go:cgo_import_dynamic libc_Getuid getuid "libc.so"
//go:cgo_import_dynamic libc_Kill kill "libc.so"
//go:cgo_import_dynamic libc_Lchown lchown "libc.so"
//go:cgo_import_dynamic libc_Link link "libc.so"
//go:cgo_import_dynamic libc___xnet_listen __xnet_listen "libsocket.so"
//go:cgo_import_dynamic libc_Lstat lstat "libc.so"
//go:cgo_import_dynamic libc_Mkdir mkdir "libc.so"
//go:cgo_import_dynamic libc_Mknod mknod "libc.so"
//go:cgo_import_dynamic libc_Nanosleep nanosleep "libc.so"
//go:cgo_import_dynamic libc_Open open "libc.so"
//go:cgo_import_dynamic libc_Pathconf pathconf "libc.so"
//go:cgo_import_dynamic libc_Pread pread "libc.so"
//go:cgo_import_dynamic libc_Pwrite pwrite "libc.so"
//go:cgo_import_dynamic libc_read read "libc.so"
//go:cgo_import_dynamic libc_Readlink readlink "libc.so"
//go:cgo_import_dynamic libc_Rename rename "libc.so"
//go:cgo_import_dynamic libc_Rmdir rmdir "libc.so"
//go:cgo_import_dynamic libc_lseek lseek "libc.so"
//go:cgo_import_dynamic libc_sendfile sendfile "libsendfile.so"
//go:cgo_import_dynamic libc_Setegid setegid "libc.so"
//go:cgo_import_dynamic libc_Seteuid seteuid "libc.so"
//go:cgo_import_dynamic libc_Setgid setgid "libc.so"
//go:cgo_import_dynamic libc_Setpgid setpgid "libc.so"
//go:cgo_import_dynamic libc_Setpriority setpriority "libc.so"
//go:cgo_import_dynamic libc_Setregid setregid "libc.so"
//go:cgo_import_dynamic libc_Setreuid setreuid "libc.so"
//go:cgo_import_dynamic libc_Setrlimit setrlimit "libc.so"
//go:cgo_import_dynamic libc_Setsid setsid "libc.so"
//go:cgo_import_dynamic libc_Setuid setuid "libc.so"
//go:cgo_import_dynamic libc_shutdown shutdown "libsocket.so"
//go:cgo_import_dynamic libc_Stat stat "libc.so"
//go:cgo_import_dynamic libc_Symlink symlink "libc.so"
//go:cgo_import_dynamic libc_Sync sync "libc.so"
//go:cgo_import_dynamic libc_Truncate truncate "libc.so"
//go:cgo_import_dynamic libc_Fsync fsync "libc.so"
//go:cgo_import_dynamic libc_Ftruncate ftruncate "libc.so"
//go:cgo_import_dynamic libc_Umask umask "libc.so"
//go:cgo_import_dynamic libc_Unlink unlink "libc.so"
//go:cgo_import_dynamic libc_utimes utimes "libc.so"
//go:cgo_import_dynamic libc___xnet_bind __xnet_bind "libsocket.so"
//go:cgo_import_dynamic libc___xnet_connect __xnet_connect "libsocket.so"
//go:cgo_import_dynamic libc_mmap mmap "libc.so"
//go:cgo_import_dynamic libc_munmap munmap "libc.so"
//go:cgo_import_dynamic libc___xnet_sendto __xnet_sendto "libsocket.so"
//go:cgo_import_dynamic libc___xnet_socket __xnet_socket "libsocket.so"
//go:cgo_import_dynamic libc___xnet_socketpair __xnet_socketpair "libsocket.so"
//go:cgo_import_dynamic libc_write write "libc.so"
//go:cgo_import_dynamic libc___xnet_getsockopt __xnet_getsockopt "libsocket.so"
//go:cgo_import_dynamic libc_getpeername getpeername "libsocket.so"
//go:cgo_import_dynamic libc_getsockname getsockname "libsocket.so"
//go:cgo_import_dynamic libc_setsockopt setsockopt "libsocket.so"
//go:cgo_import_dynamic libc_recvfrom recvfrom "libsocket.so"
//go:cgo_import_dynamic libc___xnet_recvmsg __xnet_recvmsg "libsocket.so"
//go:cgo_import_dynamic libc_getexecname getexecname "libc.so"
//go:cgo_import_dynamic libc_utimensat utimensat "libc.so"

//go:linkname libc_Getcwd libc_Getcwd
//go:linkname libc_getgroups libc_getgroups
//go:linkname libc_setgroups libc_setgroups
//go:linkname libc_fcntl libc_fcntl
//go:linkname libc_accept libc_accept
//go:linkname libc___xnet_sendmsg libc___xnet_sendmsg
//go:linkname libc_Access libc_Access
//go:linkname libc_Adjtime libc_Adjtime
//go:linkname libc_Chdir libc_Chdir
//go:linkname libc_Chmod libc_Chmod
//go:linkname libc_Chown libc_Chown
//go:linkname libc_Chroot libc_Chroot
//go:linkname libc_Close libc_Close
//go:linkname libc_Dup libc_Dup
//go:linkname libc_Fchdir libc_Fchdir
//go:linkname libc_Fchmod libc_Fchmod
//go:linkname libc_Fchown libc_Fchown
//go:linkname libc_Fpathconf libc_Fpathconf
//go:linkname libc_Fstat libc_Fstat
//go:linkname libc_Getdents libc_Getdents
//go:linkname libc_Getgid libc_Getgid
//go:linkname libc_Getpid libc_Getpid
//go:linkname libc_Geteuid libc_Geteuid
//go:linkname libc_Getegid libc_Getegid
//go:linkname libc_Getppid libc_Getppid
//go:linkname libc_Getpriority libc_Getpriority
//go:linkname libc_Getrlimit libc_Getrlimit
//go:linkname libc_Gettimeofday libc_Gettimeofday
//go:linkname libc_Getuid libc_Getuid
//go:linkname libc_Kill libc_Kill
//go:linkname libc_Lchown libc_Lchown
//go:linkname libc_Link libc_Link
//go:linkname libc___xnet_listen libc___xnet_listen
//go:linkname libc_Lstat libc_Lstat
//go:linkname libc_Mkdir libc_Mkdir
//go:linkname libc_Mknod libc_Mknod
//go:linkname libc_Nanosleep libc_Nanosleep
//go:linkname libc_Open libc_Open
//go:linkname libc_Pathconf libc_Pathconf
//go:linkname libc_Pread libc_Pread
//go:linkname libc_Pwrite libc_Pwrite
//go:linkname libc_read libc_read
//go:linkname libc_Readlink libc_Readlink
//go:linkname libc_Rename libc_Rename
//go:linkname libc_Rmdir libc_Rmdir
//go:linkname libc_lseek libc_lseek
//go:linkname libc_sendfile libc_sendfile
//go:linkname libc_Setegid libc_Setegid
//go:linkname libc_Seteuid libc_Seteuid
//go:linkname libc_Setgid libc_Setgid
//go:linkname libc_Setpgid libc_Setpgid
//go:linkname libc_Setpriority libc_Setpriority
//go:linkname libc_Setregid libc_Setregid
//go:linkname libc_Setreuid libc_Setreuid
//go:linkname libc_Setrlimit libc_Setrlimit
//go:linkname libc_Setsid libc_Setsid
//go:linkname libc_Setuid libc_Setuid
//go:linkname libc_shutdown libc_shutdown
//go:linkname libc_Stat libc_Stat
//go:linkname libc_Symlink libc_Symlink
//go:linkname libc_Sync libc_Sync
//go:linkname libc_Truncate libc_Truncate
//go:linkname libc_Fsync libc_Fsync
//go:linkname libc_Ftruncate libc_Ftruncate
//go:linkname libc_Umask libc_Umask
//go:linkname libc_Unlink libc_Unlink
//go:linkname libc_utimes libc_utimes
//go:linkname libc___xnet_bind libc___xnet_bind
//go:linkname libc___xnet_connect libc___xnet_connect
//go:linkname libc_mmap libc_mmap
//go:linkname libc_munmap libc_munmap
//go:linkname libc___xnet_sendto libc___xnet_sendto
//go:linkname libc___xnet_socket libc___xnet_socket
//go:linkname libc___xnet_socketpair libc___xnet_socketpair
//go:linkname libc_write libc_write
//go:linkname libc___xnet_getsockopt libc___xnet_getsockopt
//go:linkname libc_getpeername libc_getpeername
//go:linkname libc_getsockname libc_getsockname
//go:linkname libc_setsockopt libc_setsockopt
//go:linkname libc_recvfrom libc_recvfrom
//go:linkname libc___xnet_recvmsg libc___xnet_recvmsg
//go:linkname libc_getexecname libc_getexecname
//go:linkname libc_utimensat libc_utimensat

type libcFunc uintptr

var (
	libc_Getcwd,
	libc_getgroups,
	libc_setgroups,
	libc_fcntl,
	libc_accept,
	libc___xnet_sendmsg,
	libc_Access,
	libc_Adjtime,
	libc_Chdir,
	libc_Chmod,
	libc_Chown,
	libc_Chroot,
	libc_Close,
	libc_Dup,
	libc_Fchdir,
	libc_Fchmod,
	libc_Fchown,
	libc_Fpathconf,
	libc_Fstat,
	libc_Getdents,
	libc_Getgid,
	libc_Getpid,
	libc_Geteuid,
	libc_Getegid,
	libc_Getppid,
	libc_Getpriority,
	libc_Getrlimit,
	libc_Gettimeofday,
	libc_Getuid,
	libc_Kill,
	libc_Lchown,
	libc_Link,
	libc___xnet_listen,
	libc_Lstat,
	libc_Mkdir,
	libc_Mknod,
	libc_Nanosleep,
	libc_Open,
	libc_Pathconf,
	libc_Pread,
	libc_Pwrite,
	libc_read,
	libc_Readlink,
	libc_Rename,
	libc_Rmdir,
	libc_lseek,
	libc_sendfile,
	libc_Setegid,
	libc_Seteuid,
	libc_Setgid,
	libc_Setpgid,
	libc_Setpriority,
	libc_Setregid,
	libc_Setreuid,
	libc_Setrlimit,
	libc_Setsid,
	libc_Setuid,
	libc_shutdown,
	libc_Stat,
	libc_Symlink,
	libc_Sync,
	libc_Truncate,
	libc_Fsync,
	libc_Ftruncate,
	libc_Umask,
	libc_Unlink,
	libc_utimes,
	libc___xnet_bind,
	libc___xnet_connect,
	libc_mmap,
	libc_munmap,
	libc___xnet_sendto,
	libc___xnet_socket,
	libc___xnet_socketpair,
	libc_write,
	libc___xnet_getsockopt,
	libc_getpeername,
	libc_getsockname,
	libc_setsockopt,
	libc_recvfrom,
	libc___xnet_recvmsg,
	libc_getexecname,
	libc_utimensat libcFunc
)

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getcwd(buf []byte) (n int, err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Getcwd)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getgroups(ngid int, gid *_Gid_t) (n int, err error) {
	r0, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_getgroups)), 2, uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func setgroups(ngid int, gid *_Gid_t) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_setgroups)), 2, uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func fcntl(fd int, cmd int, arg int) (val int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_fcntl)), 3, uintptr(fd), uintptr(cmd), uintptr(arg), 0, 0, 0)
	val = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_accept)), 3, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func sendmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc___xnet_sendmsg)), 3, uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Access(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Access)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Adjtime(delta *Timeval, olddelta *Timeval) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Adjtime)), 2, uintptr(unsafe.Pointer(delta)), uintptr(unsafe.Pointer(olddelta)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Chdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Chdir)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Chmod(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Chmod)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Chown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Chown)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Chroot(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Chroot)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Close(fd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Close)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Dup(fd int) (nfd int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Dup)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	nfd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchdir(fd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Fchdir)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchmod(fd int, mode uint32) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Fchmod)), 2, uintptr(fd), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchown(fd int, uid int, gid int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Fchown)), 3, uintptr(fd), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fpathconf(fd int, name int) (val int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Fpathconf)), 2, uintptr(fd), uintptr(name), 0, 0, 0, 0)
	val = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fstat(fd int, stat *Stat_t) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Fstat)), 2, uintptr(fd), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getdents(fd int, buf []byte, basep *uintptr) (n int, err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Getdents)), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(unsafe.Pointer(basep)), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getgid() (gid int) {
	r0, _, _ := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Getgid)), 0, 0, 0, 0, 0, 0, 0)
	gid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getpid() (pid int) {
	r0, _, _ := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Getpid)), 0, 0, 0, 0, 0, 0, 0)
	pid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Geteuid() (euid int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&libc_Geteuid)), 0, 0, 0, 0, 0, 0, 0)
	euid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getegid() (egid int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&libc_Getegid)), 0, 0, 0, 0, 0, 0, 0)
	egid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getppid() (ppid int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&libc_Getppid)), 0, 0, 0, 0, 0, 0, 0)
	ppid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getpriority(which int, who int) (n int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Getpriority)), 2, uintptr(which), uintptr(who), 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getrlimit(which int, lim *Rlimit) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Getrlimit)), 2, uintptr(which), uintptr(unsafe.Pointer(lim)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Gettimeofday(tv *Timeval) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Gettimeofday)), 1, uintptr(unsafe.Pointer(tv)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getuid() (uid int) {
	r0, _, _ := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Getuid)), 0, 0, 0, 0, 0, 0, 0)
	uid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Kill(pid int, signum Signal) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Kill)), 2, uintptr(pid), uintptr(signum), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Lchown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Lchown)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Link(path string, link string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(link)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Link)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Listen(s int, backlog int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc___xnet_listen)), 2, uintptr(s), uintptr(backlog), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Lstat(path string, stat *Stat_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Lstat)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mkdir(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Mkdir)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mknod(path string, mode uint32, dev int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Mknod)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(dev), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Nanosleep(time *Timespec, leftover *Timespec) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Nanosleep)), 2, uintptr(unsafe.Pointer(time)), uintptr(unsafe.Pointer(leftover)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Open(path string, mode int, perm uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Open)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(perm), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Pathconf(path string, name int) (val int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Pathconf)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(name), 0, 0, 0, 0)
	val = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Pread(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Pread)), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Pwrite(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Pwrite)), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func read(fd int, p []byte) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_read)), 3, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Readlink(path string, buf []byte) (n int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	if len(buf) > 0 {
		_p1 = &buf[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Readlink)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), uintptr(len(buf)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Rename(from string, to string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(from)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(to)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Rename)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Rmdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Rmdir)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Seek(fd int, offset int64, whence int) (newoffset int64, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_lseek)), 3, uintptr(fd), uintptr(offset), uintptr(whence), 0, 0, 0)
	newoffset = int64(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_sendfile)), 4, uintptr(outfd), uintptr(infd), uintptr(unsafe.Pointer(offset)), uintptr(count), 0, 0)
	written = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setegid(egid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Setegid)), 1, uintptr(egid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Seteuid(euid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Seteuid)), 1, uintptr(euid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setgid(gid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Setgid)), 1, uintptr(gid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setpgid(pid int, pgid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Setpgid)), 2, uintptr(pid), uintptr(pgid), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setpriority(which int, who int, prio int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Setpriority)), 3, uintptr(which), uintptr(who), uintptr(prio), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setregid(rgid int, egid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Setregid)), 2, uintptr(rgid), uintptr(egid), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setreuid(ruid int, euid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Setreuid)), 2, uintptr(ruid), uintptr(euid), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setrlimit(which int, lim *Rlimit) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Setrlimit)), 2, uintptr(which), uintptr(unsafe.Pointer(lim)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setsid() (pid int, err error) {
	r0, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Setsid)), 0, 0, 0, 0, 0, 0, 0)
	pid = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setuid(uid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Setuid)), 1, uintptr(uid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Shutdown(s int, how int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_shutdown)), 2, uintptr(s), uintptr(how), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Stat(path string, stat *Stat_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Stat)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Symlink(path string, link string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(link)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Symlink)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Sync() (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Sync)), 0, 0, 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Truncate(path string, length int64) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Truncate)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fsync(fd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Fsync)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Ftruncate(fd int, length int64) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Ftruncate)), 2, uintptr(fd), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Umask(newmask int) (oldmask int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&libc_Umask)), 1, uintptr(newmask), 0, 0, 0, 0, 0)
	oldmask = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Unlink(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Unlink)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func utimes(path string, times *[2]Timeval) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_utimes)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(times)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc___xnet_bind)), 3, uintptr(s), uintptr(addr), uintptr(addrlen), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc___xnet_connect)), 3, uintptr(s), uintptr(addr), uintptr(addrlen), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func mmap(addr uintptr, length uintptr, prot int, flag int, fd int, pos int64) (ret uintptr, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_mmap)), 6, uintptr(addr), uintptr(length), uintptr(prot), uintptr(flag), uintptr(fd), uintptr(pos))
	ret = uintptr(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func munmap(addr uintptr, length uintptr) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_munmap)), 2, uintptr(addr), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc___xnet_sendto)), 6, uintptr(s), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(flags), uintptr(to), uintptr(addrlen))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func socket(domain int, typ int, proto int) (fd int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc___xnet_socket)), 3, uintptr(domain), uintptr(typ), uintptr(proto), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func socketpair(domain int, typ int, proto int, fd *[2]int32) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc___xnet_socketpair)), 4, uintptr(domain), uintptr(typ), uintptr(proto), uintptr(unsafe.Pointer(fd)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func write(fd int, p []byte) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_write)), 3, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc___xnet_getsockopt)), 5, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(unsafe.Pointer(vallen)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_getpeername)), 3, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_getsockname)), 3, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_setsockopt)), 5, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(vallen), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_recvfrom)), 6, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(flags), uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(fromlen)))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func recvmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc___xnet_recvmsg)), 3, uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getexecname() (path unsafe.Pointer, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_getexecname)), 0, 0, 0, 0, 0, 0, 0)
	path = unsafe.Pointer(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func utimensat(dirfd int, path string, times *[2]Timespec, flag int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_utimensat)), 4, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(times)), uintptr(flag), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
