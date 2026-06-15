package transform

import (
	"fmt"
)

func Reverse(ppwd *[]byte) {
	if ppwd == nil {
		return
	}

	fmt.Println("Reversing password... ")
	for i, j := 0, len(*ppwd)-1; i < j; i, j = i+1, j-1 {
		(*ppwd)[i], (*ppwd)[j] = (*ppwd)[j], (*ppwd)[i]
	}
	fmt.Println("Reversed password:", string(*ppwd))
	fmt.Print("Press Enter to continue... ")
	fmt.Scanln()
}
