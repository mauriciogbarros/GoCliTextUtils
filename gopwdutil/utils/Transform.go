package utils

import (
	"fmt"
	"gopwdutil/transform"
	"gopwdutil/ui"
)

func Transform(ppwd *[]byte) {
	if ppwd == nil {
		return
	}

	returnToMain := false
	choice := 0

	for !returnToMain {
		choice = ui.TransformMenu()
		switch choice {
		case 0:
			returnToMain = true
		case 1:
			transform.Reverse(ppwd)
		case 2:
			transform.Scramble(ppwd)
		case 3:
			transform.Hash(ppwd)
		case 4:
			transform.Salt(ppwd)
		}
		fmt.Println()
	}
}
