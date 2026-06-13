package tools

func ContainsRune(r_value rune, ref_runes []rune) bool {
	for _, r_ref := range ref_runes {
		if r_value == r_ref { return true }
	}
	return false
}
