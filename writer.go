package shout

/*
#cgo LDFLAGS: -lshout
#include <stdlib.h>
#include <shout/shout.h>
*/
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

// Config is a shout configuration struct
type Config struct {
	Host     string
	Port     uint16
	User     string
	Password string
	Mount    string
	Proto    Protocol
	Format   StreamFormat
}

// Writer represents a Writer interface to libshout streaming structure
type Writer struct {
	cshout *C.struct_shout
}

// Connect creates a connection according to configuration and returns
// a Writer instance
func Connect(cfg *Config) (*Writer, error) {
	var cstr *C.char
	shout := C.shout_new()
	if shout == nil {
		return nil, fmt.Errorf("error allocating shout_t structure")
	}

	// setting hostname
	cstr = C.CString(cfg.Host)
	C.shout_set_host(shout, cstr)
	C.free(unsafe.Pointer(cstr))

	// setting port
	C.shout_set_port(shout, C.ushort(cfg.Port))

	// setting source user
	cstr = C.CString(cfg.User)
	C.shout_set_user(shout, cstr)
	C.free(unsafe.Pointer(cstr))

	// setting source password
	cstr = C.CString(cfg.Password)
	C.shout_set_password(shout, cstr)
	C.free(unsafe.Pointer(cstr))

	// setting mountpoint
	cstr = C.CString(cfg.Mount)
	C.shout_set_mount(shout, cstr)
	C.free(unsafe.Pointer(cstr))

	// setting stream format
	C.shout_set_format(shout, C.uint(cfg.Format))
	// setting stream protocol
	C.shout_set_protocol(shout, C.uint(cfg.Proto))

	w := &Writer{cshout: shout}
	runtime.SetFinalizer(w, finalize)
	return w, nil
}

func (w *Writer) Write(p []byte) (n int, err error) {
	pptr := (*C.uchar)(&p[0])
	bufSize := C.size_t(len(p))
	res := int(C.shout_send(w.cshout, pptr, bufSize))
	err = convError(res)
	if err == nil {
		n = len(p)
		C.shout_sync(w.cshout)
	}
	return
}

// Close closes a shout connection
func (w *Writer) Close() error {
	res := int(C.shout_close(w.cshout))
	return convError(res)
}

func (w *Writer) getErrno() int {
	return int(C.shout_get_errno(w.cshout))
}

func finalize(w *Writer) {
	C.shout_free(w.cshout)
}
