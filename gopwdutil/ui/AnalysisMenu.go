package ui

import (
	"fmt"
	"gopwdutil/analysis"
	"gopwdutil/tools"
	"strings"
)

func Analysis(ppwd *[]rune) {
	if ppwd == nil { return }

	returnToMain := false
	choice := 0
	length := 0

	for !returnToMain {
		length = len(*ppwd)
		
		fmt.Println("======== Analysis Tools ==========")
		tools.DisplayCurrentPassword(ppwd)
		if length > 0 {
			fmt.Println("1) Character count")
			fmt.Println("2) Word count")
			fmt.Println("3) Letter count")
			fmt.Println("4) Upper count")
			fmt.Println("5) Lower count")
			fmt.Println("6) Numeric count")
			fmt.Println("7) Special count")
			fmt.Println("8) Repeated counts")
			fmt.Println("9) Strength analysis")
		}
		fmt.Println("0) Return to Main Menu")
		fmt.Println(strings.Repeat("-", 34))

		if length > 0 {
			choice = tools.GetMenuChoice(9)
		} else {
			choice = tools.GetMenuChoice(0)
		}

		if choice == 0 {
			returnToMain = true
		} else if length > 0 {
			switch choice {
				case 8: analysis.Repeated(ppwd)
				case 9: analysis.Strength(ppwd)
				default: analysis.Count(ppwd, choice)
			}
		}

		fmt.Println()
	}
}
