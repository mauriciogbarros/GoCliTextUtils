package utils

import (
	"fmt"
	"gopwdutil/analysis"
	"gopwdutil/tools"
	"gopwdutil/ui"
)

func Analysis(ppwd *[]byte) error {
	if ppwd == nil {
		return tools.Errors.NilError
	}
	returnToMain := false
	var choice int
	var err error
	for !returnToMain {
		choice, err = ui.AnalysisMenu()
		if err != nil {
			return err
		}
		switch choice {
		case 0:
			returnToMain = true
		case 8:
			analysis.Repeated(ppwd)
		case 9:
			analysis.Strength(ppwd)
		default:
			analysis.Count(ppwd, choice)
		}
		fmt.Println()
	}

	return nil
}
