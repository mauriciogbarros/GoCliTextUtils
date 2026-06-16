package transform

import (
	"fmt"
	"gopwdutil/tools"
	"strings"
)

func Pepper(ppepper *[]byte) error {
	if ppepper == nil {
		return tools.Errors.NilError
	}

	fmt.Printf("Min: %d, Max: %d\n", tools.MinPepper, tools.MaxPepper)
	fmt.Print("Enter the pepper: ")

	var newPepper []byte
	for {
		b, err := tools.Reader.ReadByte()
		if err != nil {
			fmt.Println("Error: failed to capture")
			return err
		}
		if b == '\n' || b == '\r' {
			break
		}

		newPepper = append(newPepper, b)
	}

	fmt.Println("...")
	if len(newPepper) < tools.MinPepper || len(newPepper) > tools.MaxPepper {
		fmt.Println("Error: invalid password length.")
		tools.Reset(&newPepper)
		return tools.Errors.LengthError
	}

	if strings.TrimSpace(string(newPepper)) == "" {
		fmt.Println("Error: white space only.")
		tools.Reset(&newPepper)
		return tools.Errors.WhiteSpace
	}

	tools.Reset(ppepper)
	*ppepper = append(*ppepper, newPepper...)
	tools.Reset(&newPepper)

	fmt.Println("Pepper accepted")

	return nil
}