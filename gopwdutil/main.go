package main

import (
	"fmt"
	"gopwdutil/tools"
	"gopwdutil/ui"
	"gopwdutil/utils"
)

func main() {
	var pwd tools.Password

	fmt.Println("======= CLI Password Utils =======")
	exit := false
	var choice int
	var err error
	for !exit {
		choice, err = ui.MainMenu(len(pwd.Password))
		if err != nil {
			fmt.Println(err)
		}
		if choice == 0 {
			exit = true
		} else {
			switch choice {
			case 1:
				utils.Input(&pwd.Password)
			case 2:
				utils.Analysis(&pwd.Password)
			case 3:
				utils.Transform(&pwd)
			case 4:
				utils.Output(&pwd)
			case 5:
				tools.Reset(&pwd.Password)
				tools.Reset(&pwd.Pepper)
				tools.Reset(&pwd.Salt)
				tools.Reset(&pwd.HashedKey)
			}
		}
	}

	fmt.Println("Erasing everything ....")
	// Zero out the bytes of password, pepper, salt and hashed key in memory before exit to avoid leaving sensitive data behind
	tools.Reset(&pwd.Password)
	tools.Reset(&pwd.Salt)
	tools.Reset(&pwd.Pepper)
	tools.Reset(&pwd.HashedKey)
	fmt.Println("...")
	fmt.Println("Everything erased.")
	fmt.Println("Goodbye!")
	fmt.Println()
}
