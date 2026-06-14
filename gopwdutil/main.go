package main

import (
	"fmt"
	"gopwdutil/ui"
	"gopwdutil/utils"
)

func main() {
	const minLength = 8
	const maxLength = 72

	pwd := make([]byte, 0, maxLength)
	for i := range pwd { pwd[i] = 0 }

	fmt.Println("======= CLI Password Utils =======")
	exit := false
	choice := 0
	for !exit {
		choice = ui.MainMenu(len(pwd))
		if choice == 0 {
			exit = true
		} else {
			switch choice {
				case 1: utils.Input(&pwd, minLength, maxLength)
				case 2: utils.Analysis(&pwd)
				case 3: utils.Transform(&pwd)
				case 4: utils.Output(&pwd)
			}
		}
	}

	fmt.Println("Erasing password ....")
	for i := range pwd { pwd[i] = 0}
	fmt.Println("...")
	fmt.Println("Password erased.")
	fmt.Println("Goodbye!")
	fmt.Println()
}