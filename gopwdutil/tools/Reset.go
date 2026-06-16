package tools

func Reset(word *[]byte) {
	if word == nil {
		return
	}

	// Zero out each byte before shrinking to prevent lingering sensitive data in memory
	for i := range *word {
		(*word)[i] = 0
	}

	// Shrink slice length without freeing the underlying array
	*word = (*word)[:0]
}
