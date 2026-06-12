package tools

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetMenuChoice(maxOption int) int {
	fmt.Print("Choice: ")
	line, _ := reader.ReadString('\n')
	raw := strings.TrimRight(line, "\r\n")
	choice, err := strconv.Atoi(raw)
	if err != nil {
		fmt.Printf("Error: Please enter a number, got %q.\n", line)
		return -1
	}

	if choice < 0 || choice > maxOption {
		fmt.Println("Error: Invalid option.")
		return -1
	}

	return choice
}