package transform

import (
	"fmt"
)

func Salt(ppwd *[]byte) {
	if ppwd == nil {
		return
	}

	fmt.Println("Salt password: under construction")
}
