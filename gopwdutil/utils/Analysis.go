package utils

import (
	"fmt"
	"gopwdutil/analysis"
	"gopwdutil/ui"
)

func Analysis(ppwd *[]byte) {
	if ppwd == nil { return }
	returnToMain := false
	choice := 0
	for !returnToMain {
		choice = ui.AnalysisMenu()
			switch choice {
			case 0: returnToMain = true
			case 8: analysis.Repeated(ppwd)
			case 9: analysis.Strength(ppwd)
			default: analysis.Count(ppwd, choice)
		}
		fmt.Println()
	}
}