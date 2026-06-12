package main

import (
	"fmt"
	"gopwdutil/ui"
)

func main() {
	pwd := ""
	defer func() { pwd = "" }()

	fmt.Println("===== Go Password Utils =====")
	ui.ShowMainMenu(&pwd)
	fmt.Println("Password erased... Goodbye!")
}