package utils

import (
	"fmt"
	"gopwdutil/output"
	"gopwdutil/tools"
	"gopwdutil/ui"
)

func Output(ppwd *tools.Password) error {
	if ppwd == nil {
		return tools.Errors.NilError
	}

	returnToMain := false
	var choice int
	var err error

	for !returnToMain {
		choice, err = ui.OutputMenu()
		if err != nil {
			return tools.Errors.NilError
		}

		switch choice {
		case 0:
			returnToMain = true
		case 1:
			output.ToScreen(ppwd)
		case 2:
			output.ToClipboard(ppwd)
		case 3:
			output.ToFile(ppwd)
		}
		fmt.Println()
	}

	return nil
}
