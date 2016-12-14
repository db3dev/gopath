// +build windows

package util

import "strings"

func CheckForProfile() bool {
	return createProfile()
}

// ParseEnvForGoPath Extracts the GOPATH out of the current environment
func (w *ProfileWriter) ParseEnvForGoPath() string {
	for i := range w.env {
		if strings.Contains(strings.ToLower(w.env[i]), "gopath=") {
			return w.env[i]
		}
	}

	// GOPATH does not exist in environment
	return ""
}

// ParseProfileForAlias Creates a slice of alias from a profile
func (w *ProfileWriter) ParseProfileForAlias(alias string) []string {
	return nil
}
