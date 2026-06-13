package output

import (
	"fmt"
)

func ToScreen(ppwd *[]rune) {
	if ppwd == nil { return }
	
	fmt.Println("Password:", *ppwd)
}