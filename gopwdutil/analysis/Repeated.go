package analysis

import (
	"fmt"
	"gopwdutil/tools"
)

func Repeated(ppwd *[]byte) error {
	if ppwd == nil {
		return tools.Errors.NilError
	}

	fmt.Println("======> Repeated:")

	repeated, err := getRepeated(ppwd)
	if err != nil {
		return err
	}
	
	if len(repeated) == 0 {
		fmt.Println("        No repeated characters.")
	} else {
		for r, n := range repeated {
			if n > 1 {
				fmt.Printf("        '%c': %d times\n", r, n)
			}
		}
	}

	fmt.Print("        Press Enter to continue... ")
	fmt.Scanln()

	return nil
}

func getRepeated(ppwd *[]byte) (map[byte]int, error) {
	if ppwd == nil {
		return nil, tools.Errors.NilError
	}

	repeated := map[byte]int{}
	for _, b := range *ppwd {
		if b != 0 {
			repeated[b]++
		}
	}

	for b, n := range repeated {
		if n == 1 {
			delete(repeated, b)
		}
	}

	return repeated, nil
}
