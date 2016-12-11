// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package util

import (
	"errors"
	"fmt"
	"os"
)

func CheckForProfile() bool {
	if _, err := os.Stat("~/.gopath"); err != nil {
		// If profile does not exist create it
		return createProfile()
	}

	return true
}

func createProfile() bool {
	fmt.Println(os.Getenv("GOPATH"))

	fmt.Println(errors.New("Not Implemented").Error())

	return false
}

func parsePaths() {

}
