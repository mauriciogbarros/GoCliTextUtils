package analysis

import (
	"fmt"
	"gopwdutil/tools"
	"strings"
)

func WordCount(ppwd *string) (int) {
	if ppwd == nil { return -1}

	if !tools.IsValidString(ppwd) {
		fmt.Println("Error: invalid string provided.")
		return -1
	}
	
	var count int = len(strings.Fields(*ppwd))

	fmt.Printf("Count: %d words", count)
	return count
}