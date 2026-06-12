package ui

import (
	"fmt"
	"gopwdutil/analysis"
	"gopwdutil/tools"
)

func ShowAnalysisMenu(ppwd *string) {
	if ppwd == nil { return }

	returnToMain := false
	choice := 0

	for !returnToMain {
		fmt.Println("======= Password Analysis ========")
		tools.DisplayCurrentPassword(ppwd)
		fmt.Println("1) Character count")
		fmt.Println("2) Word count")
		fmt.Println("3) Letter count")
		fmt.Println("4) Upper count")
		fmt.Println("5) Lower count")
		fmt.Println("6) Numeric count")
		fmt.Println("7) Special count")
		fmt.Println("8) Repeated counts")
		fmt.Println("9) Strength analysis")
		fmt.Println("0) Return to Main Menu")
		choice = tools.GetMenuChoice(9)
		fmt.Println()

		switch choice {
			case -1:
			case 0: returnToMain = true
			case 1: analysis.CharCount(ppwd)
			case 2: analysis.WordCount(ppwd)
			case 3: analysis.LetterCount(ppwd)
			case 4: analysis.UpperCount(ppwd)
			case 5: analysis.LowerCount(ppwd)
			case 6: analysis.NumericCount(ppwd)
			case 7: analysis.SpecialCount(ppwd)
			case 8: analysis.RepeatedCount(ppwd)
			case 9: analysis.StrengthAnalysis(ppwd)
		}
	}

}