package main

import (
	"os"

	"fmt"

	"github.com/db3dev/gopath/commands"
	"github.com/db3dev/gopath/util"
)

const (
	fileSeperator = os.PathSeparator
	pathSeperator = os.PathListSeparator
)

func main() {

	if util.CheckForProfile() {

		if len(os.Args) > 1 {
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
		} else {
			fmt.Println("GOPATH: Completed!")
		}

	}
}
