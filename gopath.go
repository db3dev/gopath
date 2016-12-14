package main

import (
	"os"

	"fmt"

	"github.com/db3dev/gopath/commands"
	"github.com/db3dev/gopath/util"
)

func main() {

	if util.CheckForProfile() {

		if len(os.Args) > 1 {
			switch os.Args[1] {
			case "add":
				commands.Add()

			case "rem":
				commands.Rem()

			case "del":
				commands.Del()

			case "edit":
				commands.Edit()

			case "-p":
				commands.Primary(os.Args)

			default:
				commands.GoWrap()
			}
		} else {
			fmt.Println("GOPATH: Completed!")
		}
	}
}
