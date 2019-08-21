package shout

import (
	"testing"
)

func TestVersion(t *testing.T) {
	ver := GetVersion()
	if ver.Version == "" {
		t.Error("version should not be empty")
	}

	if ver.Major < 1 {
		t.Error("major should not less than 1")
	}
}
