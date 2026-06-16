package utils

import (
	"fmt"
	"gopwdutil/tools"
	"gopwdutil/transform"
	"gopwdutil/ui"
)

func Transform(ppwd *tools.Password) error {
	if ppwd == nil {
		return tools.Errors.NilError
	}

	returnToMain := false
	var choice int
	var err error

	for !returnToMain {
		choice, err = ui.TransformMenu()
		if err != nil {
			return err
		}
		switch choice {
		case 0:
			returnToMain = true
		case 1:
			transform.Reverse(&ppwd.Password)
		case 2:
			transform.Scramble(&ppwd.Password)
		case 3:
			transform.Pepper(&ppwd.Pepper)
		case 4:
			transform.Salt(&ppwd.Salt)
		case 5:
			transform.Hash(ppwd)
		}
		fmt.Println()
	}

	return nil
}
