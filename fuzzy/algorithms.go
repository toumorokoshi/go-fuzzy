package fuzzy

import "strings"

// all algorithms ignore case unless specified otherwise

// returns the levenshtein distance of two strings
func Levenshtein(a, b string) int {
	distanceMatrix := make([][]int, len(a)+1)
	for i := range distanceMatrix {
		distanceMatrix[i] = make([]int, len(b)+1)
	}
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
				min = distanceMatrix[posA][posB]
			} else {
				del := distanceMatrix[posA][posB+1] + 1 // requires deletion
				add := distanceMatrix[posA+1][posB] + 1 // an addition
				sub := distanceMatrix[posA][posB] + 1   // a substitution
				if del <= add && del <= sub {
					min = del
				} else if add <= sub {
					min = add
				} else {
					min = sub
				}
			}
			distanceMatrix[posA+1][posB+1] = min
		}
	}
	return distanceMatrix[len(a)][len(b)]
}

// returns true if the characters in a 
// appear in the same sequence as b
func SequenceMatch(a, b string) bool {
	return SequenceMatchCaseSensitive(strings.ToLower(a), strings.ToLower(b))
}

func SequenceMatchCaseSensitive(a, b string) bool {
	index := 0
	runes := []rune(a)
	rune_length := len(runes)
	for _, c := range b {
		if index < rune_length {
			if runes[index] == c {
				index++
			}
		} else {
			return true
		}
	}
	return false
}
