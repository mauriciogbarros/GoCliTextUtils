package tools

import "strings"

func IsValidString(ppwd *string) bool {
	if ppwd == nil { return false }

	return strings.TrimSpace(*ppwd) != ""
}