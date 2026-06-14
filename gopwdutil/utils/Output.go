package utils

import (
	"fmt"
	"gopwdutil/output"
	"gopwdutil/ui"
)

func Output(ppwd *[]byte) {
	if ppwd == nil { return }

	returnToMain := false
	choice := 0

	for !returnToMain {
		choice = ui.OutputMenu()

		switch choice {
			case 0: returnToMain = true
			case 1: output.ToScreen(ppwd)
			case 2: output.ToClipboard(ppwd)
			case 3: output.ToFile(ppwd)
		}
		fmt.Println()
	}
}