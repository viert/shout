package shout

/*
#cgo LDFLAGS: -lshout
#include <stdlib.h>
#include <shout/shout.h>
*/
import "C"

func init() {
	C.shout_init()
}

// ShutDown shuts down the shout library, deallocating any global storage. Don't call
// anything afterwards
func ShutDown() {
	C.shout_shutdown()
}
