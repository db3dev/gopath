package commands

import "fmt"
import "github.com/db3dev/gopath/util"

func Primary(args []string) bool {
	fmt.Println("Primary")

	util.ProfileWriter{env: []string{"test"}, profile: nil}

	return false
}
