package analysis

import (
	"errors"
	"fmt"
	"gopwdutil/tools"
	"strings"
)

func Count(ppwd *[]byte, choice int) error {
	if ppwd == nil {
		return tools.Errors.NilError
	}

	count, err := getCount(ppwd, choice)
	if err != nil {
		return err
	}

	countType := getCountType(choice)
	countUnit := getCountUnit(count, choice)

	fmt.Printf("======> %s count: %d %s\n", countType, count, countUnit)
	fmt.Print("        Press Enter to continue... ")
	fmt.Scanln()

	return nil
}

func getCountType(choice int) string {
	countType := ""
	switch choice {
	case 1:
		countType = "Character"
	case 2:
		countType = "Word"
	case 3:
		countType = "Letter character"
	case 4:
		countType = "Uppercase character"
	case 5:
		countType = "Lowercase character"
	case 6:
		countType = "Numeric character"
	case 7:
		countType = "Special character"
	}

	return countType
}

func getCount(ppwd *[]byte, choice int) (int, error) {
	if ppwd == nil {
		return 0, tools.Errors.NilError
	}
	var count int
	var err error

	switch choice {
	case 1:
		count, err = countChar(ppwd)
	case 2:
		count, err = countWord(ppwd)
	case 3:
		count, err = countLetter(ppwd)
	case 4:
		count, err = countUpper(ppwd)
	case 5:
		count, err = countLower(ppwd)
	case 6:
		count, err = countNumeric(ppwd)
	case 7:
		count, err = countSpecial(ppwd)
	}

	return count, err
}

func getCountUnit(count int, choice int) string {
	unit := ""
	if count > 0 {
		switch choice {
		case 1:
			unit = "character"
		case 2:
			unit = "word"
		case 3:
			unit = "letter"
		case 4:
			unit = "uppercase letter"
		case 5:
			unit = "lowercase letter"
		case 6:
			unit = "numeric character"
		case 7:
			unit = "special character"
		}

		if count > 1 {
			unit += "s"
		}
	}

	return unit
}

func countChar(ppwd *[]byte) (int, error) {
	if ppwd == nil {
		return 0, errors.New("password is nil")
	}
	return len(strings.ReplaceAll(string(*ppwd), " ", "")), nil
}

func countWord(ppwd *[]byte) (int, error) {
	if ppwd == nil {
		return 0, tools.Errors.NilError
	}

	return len(strings.Fields(string(*ppwd))), nil
}

func countLetter(ppwd *[]byte) (int, error) {
	if ppwd == nil {
		return 0, tools.Errors.NilError
	}

	var letters = append(tools.LowerChar, tools.UpperChar...)

	count := 0
	for i := range *ppwd {
		if tools.ContainsByte((*ppwd)[i], letters) {
			count += 1
		}
	}

	return count, nil
}

func countUpper(ppwd *[]byte) (int, error) {
	if ppwd == nil {
		return 0, tools.Errors.NilError
	}

	count := 0
	for i := range *ppwd {
		if tools.ContainsByte((*ppwd)[i], tools.UpperChar) {
			count += 1
		}
	}

	return count, nil
}

func countLower(ppwd *[]byte) (int, error) {
	if ppwd == nil {
		return 0, tools.Errors.NilError
	}

	count := 0
	for i := range *ppwd {
		if tools.ContainsByte((*ppwd)[i], tools.LowerChar) {
			count += 1
		}
	}

	return count, nil
}

func countNumeric(ppwd *[]byte) (int, error) {
	if ppwd == nil {
		return 0, tools.Errors.NilError
	}

	count := 0
	for i := range *ppwd {
		if tools.ContainsByte((*ppwd)[i], tools.NumericChar) {
			count += 1
		}
	}

	return count, nil
}

func countSpecial(ppwd *[]byte) (int, error) {
	if ppwd == nil {
		return 0, tools.Errors.NilError
	}

	count := 0
	for i := range *ppwd {
		if tools.ContainsByte((*ppwd)[i], tools.SpecialChar) {
			count += 1
		}
	}

	return count, nil
}
