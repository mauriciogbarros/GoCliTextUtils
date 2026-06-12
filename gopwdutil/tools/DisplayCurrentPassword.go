package tools

import (
	"fmt"
)

func DisplayCurrentPassword(ppwd *string) {
	if ppwd == nil { return }

	fmt.Print("Current password: ")
	if *ppwd == "" {
		fmt.Println("<Empty>")
	} else {
		fmt.Println(*ppwd)
	}
}