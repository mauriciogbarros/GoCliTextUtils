package output

import (
	"fmt"
	"gopwdutil/tools"
	"strings"

	"golang.design/x/clipboard"
)

func ToClipboard(ppwd *tools.Password) error {
	if ppwd == nil {
		return tools.Errors.NilError
	}

	fmt.Println("What would you like to output:")
	fmt.Println("1) Password")
	fmt.Println("2) Pepper")
	fmt.Println("3) Salt")
	fmt.Println("4) Hashed Key")
	fmt.Println(strings.Repeat("-", 34))
	fmt.Print("Choice: ")
	choice, err := tools.GetChoice(4)
	if err != nil {
		return err
	}

	fmt.Println("Exporting to clipboard")
	fmt.Println("...")
	switch choice {
	case 1:
		clipboard.Write(clipboard.FmtText, ppwd.Password)
		fmt.Println("Password exported to clipboard")
	case 2:
		clipboard.Write(clipboard.FmtText, ppwd.Pepper)
		fmt.Println("Pepper exported to clipboard")
	case 3:
		clipboard.Write(clipboard.FmtText, ppwd.Salt)
		fmt.Println("Salt exported to clipboard")
	case 4: 
		clipboard.Write(clipboard.FmtText, ppwd.HashedKey)
		fmt.Println("Hashed key exported to clipboard")
	}

	return nil
}
