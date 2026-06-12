package output

import (
	"fmt"
)

func ToScreen(ppwd *string) {
	if ppwd == nil { return }
	
	fmt.Println("Password:", *ppwd)
}