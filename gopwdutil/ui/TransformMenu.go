package ui

import (
	"fmt"
	"gopwdutil/tools"
	"gopwdutil/transform"
	"strings"
)

func Transform(ppwd *[]rune) {
	if ppwd == nil { return }

	returnToMain := false
	choice := -1
	length := 0

	for !returnToMain {
		length = len(*ppwd)
		
		fmt.Println("====== Transformation Tools ======")
		tools.DisplayCurrentPassword(ppwd)
		if length > 0 {
			fmt.Println("1) Reverse password")
			fmt.Println("2) Randomize password")
			fmt.Println("3) Hash password")
			fmt.Println("4) Salt password")
		}
		fmt.Println("0) Return to Main Menu")
		fmt.Println(strings.Repeat("-", 34))

		if length > 0 {
			choice = tools.GetMenuChoice(4)
		} else {
			choice = tools.GetMenuChoice(0)
		}
		fmt.Println()

		if length > 0 {
			switch choice {
				case 0: returnToMain = true
				case 1: transform.Reverse(ppwd)
				case 2: transform.Randomize(ppwd)
				case 3: transform.Hash(ppwd)
				case 4: transform.Salt(ppwd)
				default:
			} 
		} else {
			if choice == 0 { returnToMain = true }
		}
	}
}