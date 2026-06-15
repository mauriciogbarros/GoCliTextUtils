package utils

import (
	"fmt"
	"gopwdutil/input"
	"gopwdutil/ui"
)

func Input(ppwd *[]byte, minLength int, maxLength int) {
	if ppwd == nil {
		return
	}

	returnToMain := false
	choice := 0
	for !returnToMain {
		choice = ui.InputMenu()
		switch choice {
		case 0:
			returnToMain = true
		case 1:
			input.Manual(ppwd, minLength, maxLength)
		case 2:
			input.FromClipboard(ppwd, minLength, maxLength)
		case 3:
			input.Random(ppwd, minLength, maxLength)
		}
		fmt.Println()
	}
}
