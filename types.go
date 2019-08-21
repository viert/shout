package shout

/*
#cgo LDFLAGS: -lshout
#include <stdlib.h>
#include <shout/shout.h>
*/
import "C"
import "fmt"

// Version represents the libshout version struct
type Version struct {
	Version string
	Major   int
	Minor   int
	Patch   int
}

// Error is an libshout integer error type
type Error int

// Error constants
const (
	ShoutErrorSuccess     Error = 0
	ShoutErrorInsane      Error = -1
	ShoutErrorNoConnect   Error = -2
	ShoutErrorNoLogin     Error = -3
	ShoutErrorSocket      Error = -4
	ShoutErrorMalloc      Error = -5
	ShoutErrorMetadata    Error = -6
	ShoutErrorConnected   Error = -7
	ShoutErrorUnconnected Error = -8
	ShoutErrorUnsupported Error = -9
	ShoutErrorBusy        Error = -10
	ShoutErrorNoTLS       Error = -11
	ShoutErrorTLSBadCert  Error = -12
	ShoutErrorRetry       Error = -13
)

func (e Error) Error() string {
	switch e {
	case ShoutErrorSuccess:
		return "success"
	case ShoutErrorInsane:
		return "insane error"
	case ShoutErrorNoConnect:
		return "connect error"
	case ShoutErrorNoLogin:
		return "login error"
	case ShoutErrorSocket:
		return "socket error"
	case ShoutErrorMalloc:
		return "memory allocation error"
	case ShoutErrorMetadata:
		return "metadata error"
	case ShoutErrorConnected:
		return "connected"
	case ShoutErrorUnconnected:
		return "not connected"
	case ShoutErrorUnsupported:
		return "not supported"
	case ShoutErrorBusy:
		return "non-blocking io busy"
	case ShoutErrorNoTLS:
		return "no tls error"
	case ShoutErrorTLSBadCert:
		return "bad certificate"
	case ShoutErrorRetry:
		return "retry"
	default:
		return fmt.Sprintf("unknown error code %d", e)
	}
}

// GetVersion returns libshout version
func GetVersion() *Version {
	var major, minor, patch C.int
	csv := C.shout_version(
		&major,
		&minor,
		&patch,
	)

	result := &Version{
		Version: C.GoString(csv),
		Major:   int(major),
		Minor:   int(minor),
		Patch:   int(patch),
	}

	return result
}

// StreamFormat type represents a stream format
type StreamFormat int

// Format constants
const (
	ShoutFormatOGG       StreamFormat = 0
	ShoutFormatVorbis    StreamFormat = 0
	ShoutFormatMP3       StreamFormat = 1
	ShoutFormatWEBM      StreamFormat = 2
	ShoutFormatWEBMAudio StreamFormat = 3
)

// Protocol type represents a stream protocol type
type Protocol int

// Protocol constants
const (
	ProtocolHTTP       Protocol = 0
	ProtocolXAudioCast Protocol = 1
	ProtocolICY        Protocol = 2
	ProtocolRoarAudio  Protocol = 3
)

func convError(errCode int) error {
	if errCode == int(ShoutErrorSuccess) {
		return nil
	}
	return Error(errCode)
}
