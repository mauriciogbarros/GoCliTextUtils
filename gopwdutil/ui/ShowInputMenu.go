package ui

import (
	"fmt"
	"gopwdutil/input"
	"gopwdutil/tools"
)

func ShowInputMenu(ppwd *string) {
	if ppwd == nil { return }

	returnToMain := false
	choice := 0
	
	for !returnToMain {
		fmt.Println("======== Input Menu =========")
		tools.DisplayCurrentPassword(ppwd)
		fmt.Println("1) Manual input")
		fmt.Println("2) Load from clipboard")
		fmt.Println("3) Generate random string")
		fmt.Println("0) Return to Main Menu")
		choice = tools.GetMenuChoice(3)
		fmt.Println()

		switch choice {
			case -1: 
			case 0: returnToMain = true
			case 1: input.Manual(ppwd)
			case 2: input.FromClipboard(ppwd)
			case 3: input.GenerateRandom(ppwd)
		}
	}
}