package input

import (
	"fmt"
	"gopwdutil/tools"
	"strings"
)

func Manual(ppwd *[]byte) error {
	if ppwd == nil {
		return tools.Errors.NilError
	}

	fmt.Println("========== Manual Input ==========")
	fmt.Printf("From %d to %d characters long\n", tools.MinPassword, tools.MaxPassword)
	fmt.Print("New password: ")

	var newPwd []byte
	for {
		b, err := tools.Reader.ReadByte()
		if err != nil {
			fmt.Println("Error: failed to capture")
			return tools.Errors.CaptureError
		}
		if b == '\n' || b == '\r' {
			break
		}

		newPwd = append(newPwd, b)
	}

	if len(newPwd) < tools.MinPassword || len(newPwd) > tools.MaxPassword {
		fmt.Println("Error: invalid password length.")
		tools.Reset(&newPwd)
		return tools.Errors.LengthError
	}

	if strings.TrimSpace(string(newPwd)) == "" {
		fmt.Println("Error: white space only.")
		tools.Reset(&newPwd)
		return tools.Errors.WhiteSpace
	}

	tools.Reset(ppwd)

	*ppwd = append(*ppwd, newPwd...)

	tools.Reset(&newPwd)

	return nil
}
