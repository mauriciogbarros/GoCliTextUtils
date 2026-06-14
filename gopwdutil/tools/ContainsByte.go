package tools

func ContainsByte(b_value byte, ref_bytes []byte) bool {
	for _, b_ref := range ref_bytes {
		if b_value == b_ref { return true }
	}
	return false
}
