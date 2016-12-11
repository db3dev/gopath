// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package gopath

import (
	"errors"
	"fmt"
	"os"
)

func checkForProfile() bool {
	if _, err := os.Stat("~/.gopath"); err == nil {
		return false
	}

	return true
}

func createProfile() bool {
	fmt.Println(errors.New("Not Implemented").Error())

	return false
}

func parsePaths() {

}
