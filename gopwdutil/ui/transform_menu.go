package ui

import (
	"fmt"
	"gopwdutil/tools"
	"strings"
)

func TransformMenu() (int, error) {
	fmt.Println("====== Transformation Tools ======")
	fmt.Println("1) Reverse password")
	fmt.Println("2) Scramble password")
	fmt.Println("3) Crack fresh pepper")
	fmt.Println("4) Sprinkle some salt")
	fmt.Println("5) Hash password")
	fmt.Println("0) Return to Main Menu")
	fmt.Println(strings.Repeat("-", 34))
	fmt.Print("Choice: ")

	return tools.GetChoice(5)
}
