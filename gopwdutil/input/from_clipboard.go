package input

import (
	"fmt"
	"gopwdutil/tools"
	"strings"

	"golang.design/x/clipboard"
)

func FromClipboard(ppwd *[]byte) error {
	if ppwd == nil {
		return tools.Errors.NilError
	}

	fmt.Println("====== Load from Clipboard =======")
	if err := clipboard.Init(); err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// Clear clipboard first to ensure we read only what the user explicitly copies next
	clipboard.Write(clipboard.FmtText, []byte(""))
	fmt.Print("Copy string to clipboard and press Enter to continue... ")
	fmt.Scanln()

	fmt.Println("Reading from clipboard...")
	data := clipboard.Read(clipboard.FmtText)

	if data == nil {
		fmt.Println("Error: clipboard is empty.")
		return tools.Errors.EmptyError
	}

	if strings.TrimSpace(string(data)) == "" {
		fmt.Println("Error: white space only.")
		clipboard.Write(clipboard.FmtText, []byte(""))
		return tools.Errors.WhiteSpace
	}

	if len(data) < tools.MinPassword || len(data) > tools.MaxPassword {
		fmt.Println("Error: invalid password length")
		clipboard.Write(clipboard.FmtText, []byte(""))
		return tools.Errors.LengthError
	}

	tools.Reset(ppwd)

	*ppwd = append(*ppwd, data...)

	clipboard.Write(clipboard.FmtText, []byte(""))
	tools.Reset(&data)

	return nil
}
