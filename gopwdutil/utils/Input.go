package utils

import (
	"fmt"
	"gopwdutil/input"
	"gopwdutil/tools"
	"gopwdutil/ui"
)

func Input(ppwd *[]byte) error {
	if ppwd == nil {
		return tools.Errors.NilError
	}

	returnToMain := false
	var choice int
	var err error
	for !returnToMain {
		choice, err = ui.InputMenu()
		if err != nil {
			return err
		}
		switch choice {
		case 0:
			returnToMain = true
		case 1:
			input.Manual(ppwd)
		case 2:
			input.FromClipboard(ppwd)
		case 3:
			input.Random(ppwd)
		}
		fmt.Println()
	}

	return nil
}
