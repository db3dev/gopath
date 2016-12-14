// +build windows

package util

import "strings"

func CheckForProfile() bool {
	return createProfile()
}

// ParseEnvForGoPath Extracts the GOPATH out of the current environment
func (w *ProfileWriter) ParseEnvForGoPath() string {
	for i := range w.Env {
		if strings.Contains(strings.ToLower(w.Env[i]), "gopath=") {
			return w.Env[i]
		}
	}

	// GOPATH does not exist in environment
	return ""
}

// ParseProfileForAlias Creates a slice of alias from a profile
func (w *ProfileWriter) ParseProfileForAlias(alias string) []string {
	return nil
}
