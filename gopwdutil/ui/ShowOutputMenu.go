package ui

import (
	"fmt"
	"gopwdutil/output"
	"gopwdutil/tools"
)

func ShowOutputMenu(ppwd *string) {
	if ppwd == nil { return }

	returnToMain := false
	choice := 0

	for !returnToMain {
		fmt.Println("======== Output Menu ========")
		fmt.Println("1) To screen")
		fmt.Println("2) To clipboard")
		fmt.Println("3) To file")
		fmt.Println("0) Return to Main Menu")
		choice = tools.GetMenuChoice(3)
		fmt.Println()

		switch choice {
			case -1: 
			case 0: returnToMain = true
			case 1: output.ToScreen(ppwd)
			case 2: output.ToClipboard(ppwd)
			case 3: output.ToFile(ppwd)
		}
	}
}