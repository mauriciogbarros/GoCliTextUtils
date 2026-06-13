package analysis

import (
	"fmt"
)

func Repeated(ppwd *[]rune) {
	fmt.Println("======> Repeated:")

	if ppwd == nil { return }
	counts := map[rune]int{}
	for _, r := range *ppwd {
		if r != 0 {
			counts[r]++
		}
	}

	for r, n := range counts {
		if n == 1 {
			delete(counts, r)
		}
	}

	if len(counts) == 0 {
		fmt.Println("        No repeated characters.")
	} else {
		for r, n := range counts {
			if n > 1 {
				fmt.Printf("        '%c': %d times\n", r, n)
			}
		}
	}
	
	fmt.Print("        Press Enter to continue... ")
	fmt.Scanln()
}