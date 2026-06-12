package analysis

import (
	"fmt"
	"gopwdutil/tools"
	"strings"
)

func WordCount(ppwd *string) (int) {
	if ppwd == nil { return -1}
	
	if tools.IsEmpty(*ppwd) {
		fmt.Println("Error: Empty string provided.")
		return -1
	}

	if tools.IsWhiteSpace(*ppwd) {
		fmt.Println("Error: Only white spaces.")
		return -1
	}

	var count int = len(strings.Fields(*ppwd))

	fmt.Printf("Count: %d words", count)
	return count
}