// +build windows

package util

import (
	"errors"
	"fmt"
)

func CheckForProfile() bool {
	return createProfile()
}

func createProfile() bool {
	fmt.Println(errors.New("Not Implemented").Error())

	return false
}

func parsePaths() {

}
