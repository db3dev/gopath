// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package util

import "strings"

// ParseEnvForGoPath Extracts the GOPATH out of the current environment
func (w *ProfileWriter) ParseEnvForGoPath() string {
	for i := range w.Env {
		if strings.Contains(w.Env[i], "GOPATH=") {
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
