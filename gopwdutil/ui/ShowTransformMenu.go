package ui

import (
	"fmt"
	"gopwdutil/tools"
	"gopwdutil/transform"
)

func ShowTransformMenu(ppwd *string) {
	if ppwd == nil { return }

	returnToMain := false
	choice := -1

	for !returnToMain {
		fmt.Println("=== Transformations Menu ====")
		tools.DisplayCurrentPassword(ppwd)
		fmt.Println("1) Reverse password")
		fmt.Println("2) Randomize password")
		fmt.Println("3) Hash password")
		fmt.Println("4) Salt password")
		fmt.Println("0) Return to Main Menu")
		choice = tools.GetMenuChoice(4)
		fmt.Println()

		switch choice {
			case -1:
			case 0: returnToMain = true
			case 1: transform.Reverse(ppwd)
			case 2: transform.Randomize(ppwd)
			case 3: transform.Hash(ppwd)
			case 4: transform.Salt(ppwd)
		}
	}
}