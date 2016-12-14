// +build windows

package util

import "testing"

func TestParsesGoPath_WhenPresentReturnsPath(t *testing.T) {
	pw := ProfileWriter{Env: []string{"x=nope/not/i", "gOPaTH=here/i/am", "Y=not/me"}, Profile: nil}

	// Parses GOPATH when present
	if out := pw.ParseEnvForGoPath(); out == "" {
		t.Error("parseEnvForGoPath did not detect presence of GOPATH=:\n" + out)
	}

}

func TestParsesGoPath_NotPresentReturnsZero(t *testing.T) {
	pw := ProfileWriter{Env: []string{"x=nope/not/i", "gorpath=here/i/am", "Y=not/me"}, Profile: nil}

	// Returns zero value when GOPATH is not present
	if out := pw.ParseEnvForGoPath(); out != "" {
		t.Error("parseEnvForGoPath did not return zero value:\n" + out)
	}
}
