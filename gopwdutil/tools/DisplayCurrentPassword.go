package tools

import "fmt"

func DisplayCurrentPassword(ppwd *[]rune) {
	if ppwd == nil { return }
	
	fmt.Print("Current password: ")
	if len(*ppwd) == 0 {
		fmt.Println("<Empty>")
	} else {
		fmt.Println(string(*ppwd))
	}
}