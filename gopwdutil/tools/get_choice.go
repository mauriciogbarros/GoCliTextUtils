package tools

import (
	"fmt"
	"strconv"
	"strings"
)

func GetChoice(maxOption int) (int, error) {
	line, _ := Reader.ReadString('\n')
	raw := strings.TrimRight(line, "\r\n")
	choice, err := strconv.Atoi(raw)
	if err != nil {
		fmt.Printf("Error: Please enter a number, got %q.\n", line)
		return 0, err
	}

	if maxOption == 0 {
		if choice != 0 {
			fmt.Println("Error: Invalid option.")
			return 0, Errors.InvalidInput
		}
	} else if choice < 0 || choice > maxOption {
		fmt.Println("Error: Invalid option.")
		return 0, Errors.InvalidInput
	}

	return choice, nil
}
