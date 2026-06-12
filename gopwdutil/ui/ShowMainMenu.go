package ui

import (
	"fmt"
	"gopwdutil/tools"
)

func ShowMainMenu() {
	exit := false
	choice := 0
	pwd := ""
	defer func() { pwd = "" }()
	
	for !exit {
		fmt.Println("========= Main Menu =========")
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
			case 1:	ShowInputMenu(&pwd)
			case 2:	ShowAnalysisMenu(&pwd)
			case 3: ShowTransformMenu(&pwd)
			case 4: ShowOutputMenu(&pwd)
		}
	}
}