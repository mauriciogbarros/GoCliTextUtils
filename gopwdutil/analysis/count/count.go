package count

import (
	"fmt"
	"strings"
)

func Chars(ppwd *string) int {
	if ppwd == nil { return -1 }
	return len(*ppwd)
}

func Words(ppwd *string) int {
	if ppwd == nil { return -1 }
	return len(strings.Fields(*ppwd))
}

func Letters(ppwd *string) int {
	if ppwd == nil { return -1 }
	count := 0

	for _, char := range(*ppwd) {
		fmt.Println(char)
	}

	return count
}