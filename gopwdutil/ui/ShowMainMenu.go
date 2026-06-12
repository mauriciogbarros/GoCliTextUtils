package ui

import (
	"fmt"
	"gopwdutil/tools"
)

func ShowMainMenu(ppwd *string) {
	exit := false
	choice := 0
	
	for !exit {
		fmt.Println("========= Main Menu =========")
		tools.DisplayCurrentPassword(ppwd)
		fmt.Println("1) Input tools")
		fmt.Println("2) Analysis tools")
		fmt.Println("3) Transformation tools")
		fmt.Println("4) Output options")
		fmt.Println("0) Exit")
		choice = tools.GetMenuChoice(4)
		fmt.Println()
		
		switch choice {
			case -1:
			case 0: exit = true
			case 1:	ShowInputMenu(ppwd)
			case 2:	ShowAnalysisMenu(ppwd)
			case 3: ShowTransformMenu(ppwd)
			case 4: ShowOutputMenu(ppwd)
		}
	}
}