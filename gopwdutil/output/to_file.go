package output

import (
	"fmt"
	"gopwdutil/tools"
	"os"
)

func ToFile(ppwd *tools.Password) error {
	if ppwd == nil {
		return tools.Errors.NilError
	}

	var content []byte
	content = append(content, []byte("Password: ")...)
	content = append(content, ppwd.Password...)
	content = append(content, '\n')
	content = append(content, []byte("Pepper: ")...)
	content = append(content, ppwd.Pepper...)
	content = append(content, '\n')
	content = append(content, []byte("Salt: ")...)
	content = append(content, ppwd.Salt...)
	content = append(content, '\n')
	content = append(content, []byte("Hashed Key: ")...)
	content = append(content, ppwd.HashedKey...)

	err := os.WriteFile("output.txt", content, 0600)
	if err != nil {
		fmt.Println("Error: error writing to file")
		return tools.Errors.InternalError
	}

	return nil
}
