package commands

import (
	"testing"

	"github.com/db3dev/gopath/util"
)

func TestReadsWorkSpacePath(t *testing.T) {
	pw := util.ProfileWriter{Env: []string{"x=nope/not/i", "GOPATH=$GOP_FIRST:GOP_SECOND:GOP_WORKSPACE", "Y=not/me"}, Profile: nil}
	args := []string{"gopath", "-p", "WORKSPACE"}

	if Primary(args) == false {
		t.Error("Failing Test")
	}
}
