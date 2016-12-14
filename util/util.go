package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
)

// GetProfileDir Returns path for current user's .gopath (profile)
func GetProfileDir() string {
	usr, _ := user.Current()
	return usr.HomeDir + "/.gopath"
}

// GetDefaultWorkspace Returns path to fallback directory for GOPATH
func GetDefaultWorkspace() string {
	usr, _ := user.Current()
	return usr.HomeDir + "/go_workspace"
}

// CheckForProfile Check if .gopath exists and attempts to create if it does not
func CheckForProfile() bool {
	if _, err := os.Stat(GetProfileDir()); err != nil {

		// Check if profile exists
		if os.IsNotExist(err) {
			// Does not exist, create profile
			result := createProfile()
			fmt.Println("GOPATH: Created New Profile")
			return result
		}
	}

	// Profile exists
	return true
}

func createProfile() bool {
	pw := ProfileWriter{env: os.Environ(), profile: nil}

	// Create profile file
	file, err := os.Create(GetProfileDir())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	// Write Profile
	w := bufio.NewWriter(file)

	w.WriteString(
		pw.ComposeProfile(),
	)

	w.Flush()

	// Success
	return true
}

// ProfileWriter Supplies methods for writing profiles and parsing elements of the profile
type ProfileWriter struct {
	env     []string
	profile []string
}

// ComposeProfile Parses environment and existing profile alias to create a profile
func (w *ProfileWriter) ComposeProfile() string {
	goPath := w.ParseEnvForGoPath()
	goAlias := w.ParseProfileForAlias("")

	var wProfile string

	// Indicate gopath is in environment
	wProfile = "export GOPATHi=1\n"

	// Iterate and publish all aliases
	for i := range goAlias {
		wProfile += "export " + goAlias[i] + "\n"
	}

	// Write in GOPATH
	if goPath == "" {
		wProfile += "export GOPATH=" + GetDefaultWorkspace() + "\n"
	} else {
		wProfile += "export " + goPath + "\n"
	}

	// Return Profile
	return wProfile
}
