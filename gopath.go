package gopath

import (
	"os"
)

const (
	fileSeperator = os.PathSeparator
	pathSeperator = os.PathListSeparator
)

func main() {

	switch args := os.Args[1:]; args[0] {
	case "add":
		add()

	case "rem":
		rem()

	case "del":
		del()

	case "edit":
		edit()

	default:
		goWrap()
	}
}
