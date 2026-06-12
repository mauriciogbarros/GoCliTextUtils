package input

import (
	"bufio"
	"fmt"
	"gopwdutil/analysis/count"
	"gopwdutil/tools"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func Manual(ppwd *string) {
	if ppwd == nil { return }
	fmt.Println("======= Manual Input ========")
	fmt.Println("From 8 to 16 characters long")
	fmt.Print("New password: ")

	line, err := reader.ReadString('\n')
	raw := strings.TrimRight(line, "\r\n")
	if err != nil {
		fmt.Println("Error: invalid input.")
		fmt.Println()
		return
	}

	if !tools.IsValidString(&raw) {
		fmt.Println("Error: empty or white space only.")
		fmt.Println()
		return
	}

	length := count.Chars(&raw)
	if length < 8 || length > 16 {
		fmt.Println("Error: invalid password.")
		fmt.Println()
		return
	}

	*ppwd = raw
	fmt.Println()
}