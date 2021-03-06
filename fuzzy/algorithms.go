package fuzzy

import "strings"
import "unicode"

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

// returns the ranking of the matching based on the number of
// characters the character matched with, and the significance of the
// character. For example, a capital is of higher significance than a
// lowercase, so the match is ranked higher. The total rank is the
// total value of the matches. If the subsequence doesn't match, the
// value is zero.
func OrderSignificance(b, a string) int {
	a = strings.ToLower(a)
	// a value slice and optimal index slice is necessary
	runes := []rune(a)
	b_runes := []rune(b)
	index := make([][]int, len(runes), len(runes))
	for i, _ := range index {
		index[i] = make([]int, len(runes), len(runes))
		index[i][0] = -1
	}
	// array of runes left to index
	for i, c := range b_runes {
		for k, r := range reverseRunes(runes) {
			j := len(runes) - k - 1
			if unicode.ToLower(c) == r {
				if (index[j][0] == -1 && (j == 0 || index[j - 1][0] != -1)) ||
					(index[j][0] != -1 && (c == unicode.ToUpper(r) || (c == r && b_runes[index[j][0]] != unicode.ToUpper(r)))) {
					 if j != 0 {
						 copy(index[j][1:], index[j - 1][:])
					 }
					 index[j][0] = i
				 }
			}
		}
	}
	rune_index := len(runes) - 1
	value := 0
	for rune_index > -1 && index[rune_index][0] != -1 {
		if runes[rune_index] == b_runes[index[len(runes) - 1][len(runes) - rune_index - 1]] {
			value += 1
		} else { // else it's a capital
			value += 2
		}
		rune_index--
	}
	return value
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
