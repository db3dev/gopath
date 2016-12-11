package main

import (
	"os"

	"github.com/db3dev/gopath/commands"
	"github.com/db3dev/gopath/util"
)

const (
	fileSeperator = os.PathSeparator
	pathSeperator = os.PathListSeparator
)

func main() {

	if util.CheckForProfile() {
		switch args := os.Args[1:]; args[0] {
		case "add":
			commands.Add()

		case "rem":
			commands.Rem()

		case "del":
			commands.Del()

		case "edit":
			commands.Edit()

		default:
			commands.GoWrap()
		}
	}
}
