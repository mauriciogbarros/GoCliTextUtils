package input

import (
	"crypto/rand"
	"fmt"
	"gopwdutil/tools"
	"math/big"
	"strconv"
	"strings"
)

func Random(ppwd *[]byte) error {
	if ppwd == nil {
		return tools.Errors.NilError
	}

	var characters = append(tools.LowerChar, tools.UpperChar...)
	characters = append(characters, tools.NumericChar...)
	characters = append(characters, tools.SpecialChar...)

	bigLenCharacters := big.NewInt(int64(len(characters)))
	
	fmt.Println("======== Generate Random =========")
	fmt.Printf("min %d, max %d characters\n", tools.MinPassword, tools.MaxPassword)
	fmt.Print("Enter desired length: ")

	line, _ := tools.Reader.ReadString('\n')
	raw := strings.TrimRight(line, "\r\n")
	length, err := strconv.Atoi(raw)
	if err != nil {
		fmt.Printf("Error: Please enter a number, got %q.\n", line)
		return tools.Errors.InvalidInput
	}

	if length < tools.MinPassword || length > tools.MaxPassword {
		fmt.Println("Error: invalid length")
		return tools.Errors.LengthError
	}

	tools.Reset(ppwd)
	for i := 0; i < length; i++ {
		j, err := rand.Int(rand.Reader, bigLenCharacters)
		if err != nil {
			fmt.Println("Error: randomizing error")
			return tools.Errors.InternalError
		}
		*ppwd = append(*ppwd, characters[j.Int64()])
	}

	return nil
}
