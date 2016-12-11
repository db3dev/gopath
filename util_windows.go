// +build windows

package gopath

import (
	"errors"
	"fmt"
)

func checkForProfile() bool {
	return false
}

func createProfile() bool {
	fmt.Println(errors.New("Not Implemented").Error())

	return false
}

func parsePaths() {

}
