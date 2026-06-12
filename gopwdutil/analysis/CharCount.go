package analysis

import (
	"fmt"
	"gopwdutil/tools"
	"strings"
)

func CharCount(ppwd *string) (int) {
	if ppwd == nil { return -1 }
	
	fmt.Println("======= Character Count ========")
	fmt.Println("* Non-whitespace characters only")

	if !tools.IsValidString(ppwd) {
		fmt.Println("Error: Invalid string provided.")
		return -1
	}

	var length int = len(strings.ReplaceAll(*ppwd, " ", ""))
	fmt.Println("Character count:", length)
	fmt.Println()
	return length
}