package tools

func Reset(ppwd *[]rune) {
	if ppwd == nil { return }
	
	for i := range *ppwd {
		(*ppwd)[i] = 0
	}
}
