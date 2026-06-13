package ui

import (
	"fmt"
	"gopwdutil/input"
	"gopwdutil/tools"
	"strings"
)

func Input(ppwd *[]rune, minLength int, maxLength int) {
	if ppwd == nil { return }
	
	returnToMain := false
	choice := 0
	
	for !returnToMain {
		fmt.Println("========== Input Tools ===========")
		tools.DisplayCurrentPassword(ppwd)
		fmt.Println("1) Manual input")
		fmt.Println("2) Load from clipboard")
		fmt.Println("3) Generate random string")
		fmt.Println("0) Return to Main Menu")
		fmt.Println(strings.Repeat("-", 34))
		choice = tools.GetMenuChoice(3)
		fmt.Println()

		switch choice {
			case -1: 
			case 0: returnToMain = true
			case 1: input.Manual(ppwd, minLength, maxLength)
			case 2: input.FromClipboard(ppwd, minLength, maxLength)
			case 3: input.Random(ppwd, minLength, maxLength)
		}
		fmt.Println()
	}
}
