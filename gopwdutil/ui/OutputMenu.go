package ui

import (
	"fmt"
	"gopwdutil/output"
	"gopwdutil/tools"
	"strings"
)

func Output(ppwd *[]rune) {
	if ppwd == nil { return }

	returnToMain := false
	choice := 0
	length := 0

	for !returnToMain {
		length = len(*ppwd)
		
		fmt.Println("========== Output Tools ==========")
		tools.DisplayCurrentPassword(ppwd)
		if length > 0 {
			fmt.Println("1) To screen")
			fmt.Println("2) To clipboard")
			fmt.Println("3) To file")
		}
		fmt.Println("0) Return to Main Menu")
		fmt.Println(strings.Repeat("-", 34))

		if length > 0 {
			choice = tools.GetMenuChoice(3)
		} else {
			choice = tools.GetMenuChoice(0)
		}
		fmt.Println()

		if length > 0 {
			switch choice {
				case 0: returnToMain = true
				case 1: output.ToScreen(ppwd)
				case 2: output.ToClipboard(ppwd)
				case 3: output.ToFile(ppwd)
				default:
			}
		} else {
			if choice == 0 { returnToMain = true }
		}
	}
}
