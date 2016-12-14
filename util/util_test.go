package util

import "testing"

func TestComposesProfile_NewProfileFromExistingGoPath(t *testing.T) {
	pw := ProfileWriter{env: []string{"x=nope/not/i", "GOPATH=~/go_workspace", "Y=not/me"}, profile: nil}
	expectedOutput := "export GOPATHi=1\n" +
		"export GOPATH=~/go_workspace\n"

	// Returns profile as string when given only the environment
	if out := pw.ComposeProfile(); out != expectedOutput {
		t.Error("composeProfile wrote unexpected new profile when GOPATH does exist:\n" + "Result:\n" + out + "\nExpected:\n" + expectedOutput)
	}
}

func TestComposesProfile_NewProfileFromNonExistingGoPath(t *testing.T) {
	pw := ProfileWriter{env: []string{"x=nope/not/i", "gopath=~/nope_space", "Y=not/me"}, profile: nil}
	expectedOutput := "export GOPATHi=1\n" +
		"export GOPATH=" + GetDefaultWorkspace() + "\n"

	// Returns profile with default workspace location for GOPATH
	if out := pw.ComposeProfile(); out != expectedOutput {
		t.Error("composeProfile wrote unexpected new profile when GOPATH does not exist:\n" + "Result:\n" + out + "\nExpected:\n" + expectedOutput)
	}
}
