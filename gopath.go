package gopath

import (
	"fmt"
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

func add() {
	fmt.Println("add")
}

func rem() {
	fmt.Println("rem")
}

func del() {
	fmt.Println("del")
}

func edit() {
	fmt.Println("edit")
}

func goWrap() {
	fmt.Println("default")
}
