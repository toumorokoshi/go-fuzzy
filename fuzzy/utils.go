package fuzzy

func reverseRunes(a []rune) []rune {
	newSlice := make([]rune, len(a))
	copy(newSlice[:], a[:])
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    newSlice[i], newSlice[j] = newSlice[j], newSlice[i]
	}
	return newSlice
}
