package transform

import (
	"fmt"
	"gopwdutil/tools"
)

func Reverse(ppwd *[]byte) error {
	if ppwd == nil {
		return tools.Errors.NilError
	}

	fmt.Println("Reversing password... ")
	for i, j := 0, len(*ppwd)-1; i < j; i, j = i+1, j-1 {
		(*ppwd)[i], (*ppwd)[j] = (*ppwd)[j], (*ppwd)[i]
	}
	fmt.Println("Reversed password:", string(*ppwd))
	fmt.Print("Press Enter to continue... ")
	fmt.Scanln()

	return nil
}
