package analysis

import (
	"fmt"
	"gopwdutil/tools"
	"strings"
)

func Count(ppwd *[]rune, choice int) {
	if ppwd == nil { return }

	count := get_count(ppwd, choice)
	if count == -1 {
		fmt.Println("Error: password is nil")
		return
	}

	count_type := get_count_type(choice)
	count_unit := get_count_unit(count, choice)

	fmt.Printf("======> %s count: %d %s\n", count_type, count, count_unit)
	fmt.Print("        Press Enter to continue... ")
	fmt.Scanln()
}

func get_count_type(choice int) string {
	count_type := ""
	switch choice {
		case 1: count_type = "Character"
		case 2: count_type = "Word"
		case 3: count_type = "Letter character"
		case 4: count_type = "Uppercase character"
		case 5: count_type = "Lowercase character"
		case 6: count_type = "Numeric character"
		case 7: count_type = "Special character"
	}

	return count_type
}

func get_count(ppwd *[]rune, choice int) int {
	count := 0
	switch choice {
		case 1: count = count_char(ppwd)
		case 2: count = count_word(ppwd)
		case 3: count = count_letter(ppwd)
		case 4: count = count_upper(ppwd)
		case 5: count = count_lower(ppwd)
		case 6: count = count_numeric(ppwd)
		case 7: count = count_special(ppwd)
	}

	return count
}

func get_count_unit(count int, choice int) string {
	unit := ""
	if count > 0 {
		switch choice {
			case 1: unit = "character"
			case 2: unit = "word"
			case 3: unit = "letter"
			case 4: unit = "uppercase letter"
			case 5: unit = "lowercase letter"
			case 6: unit = "numeric character"
			case 7: unit = "special character"
		}

		if count > 1 {
			unit += "s"
		}
	}

	return unit
}

func count_char(ppwd *[]rune) int {
	return len(strings.ReplaceAll(string(*ppwd), " ", ""))
}

func count_word(ppwd *[]rune) int {
	return len(strings.Fields(string(*ppwd)))
}

func count_letter(ppwd *[]rune) int {
	var letters = append(lowercase, uppercase...)

	count := 0
	for i := range *ppwd {
		if tools.ContainsRune((*ppwd)[i], letters) { count += 1 }
	}

	return count
}

func count_upper(ppwd *[]rune) int {
	count := 0
	for i := range *ppwd {
		if tools.ContainsRune((*ppwd)[i], uppercase) { count += 1 }
	}

	return count
}

func count_lower(ppwd *[]rune) int {
	count := 0
	for i := range *ppwd {
		if tools.ContainsRune((*ppwd)[i], lowercase) { count += 1 }
	}

	return count
}

func count_numeric(ppwd *[]rune) int {
	count := 0
	for i := range *ppwd {
		if tools.ContainsRune((*ppwd)[i], numeric) { count += 1 }
	}

	return count
}

func count_special(ppwd *[]rune) int {
	count := 0
	for i := range *ppwd {
		if tools.ContainsRune((*ppwd)[i], special) { count += 1 }
	}

	return count
}