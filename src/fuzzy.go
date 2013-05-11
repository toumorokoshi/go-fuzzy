package fuzzy

// returns the levenshtein distance of two strings
func Levenshtein(a, b String) int {
	var distanceMatrix [len(a) + 1][len(b) + 1]int
	for pos, _ := range a {
		distanceMatrix[pos+1][0] = pos + 1
	}
	for pos, _ := range b {
		distanceMatrix[0][pos+1] = pos + 1
	}
	for posA, charA := range a {
		for posB, charB := range b {
			min := 0
			if charA == charB {
				min = d[posA][posB]
			} else {
				del := d[posA][posB+1] // requires deletion
				add := d[posA+1][posB] // an addition
				sub := d[posA][posB]   // a substitution
				if del <= add && del <= sub {
					min = del
				} else if add <= sub {
					min = add
				} else {
					min = sub
				}
			}
			d[posA+1][posB+1] = min
		}
	}
}
