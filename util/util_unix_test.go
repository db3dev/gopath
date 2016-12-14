// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package util

import "testing"

func TestParsesGoPath_WhenPresentReturnsPath(t *testing.T) {
	pw := ProfileWriter{env: []string{"x=nope/not/i", "GOPATH=here/i/am", "Y=not/me"}, profile: nil}

	// Parses GOPATH when present
	if out := pw.ParseEnvForGoPath(); out == "" {
		t.Error("parseEnvForGoPath did not detect presence of GOPATH=:\n" + out)
	}

}

func TestParsesGoPath_NotPresentReturnsZero(t *testing.T) {
	pw := ProfileWriter{env: []string{"x=nope/not/i", "gopath=here/i/am", "Y=not/me"}, profile: nil}

	// Returns zero value when GOPATH is not present
	if out := pw.ParseEnvForGoPath(); out != "" {
		t.Error("parseEnvForGoPath did not return zero value:\n" + out)
	}
}
