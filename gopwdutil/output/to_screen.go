package output

import (
	"fmt"
	"gopwdutil/tools"
)

func ToScreen(ppwd *tools.Password) error {
	if ppwd == nil {
		return tools.Errors.NilError
	}

	fmt.Printf("Password: %s\n", string(ppwd.Password))
	fmt.Printf("Pepper: %s\n", string(ppwd.Pepper))
	fmt.Printf("Salt: %s\n", string(ppwd.Salt))
	fmt.Printf("Key: %s", string(ppwd.HashedKey))
	fmt.Println()

	return nil
}
