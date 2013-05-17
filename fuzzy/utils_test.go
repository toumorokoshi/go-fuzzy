package fuzzy

import "testing"

func TestReverse(t *testing.T) {
	runes := []rune{'a','b','c'}
	reverse := []rune{'c','b','a'}
	result := reverseRunes(runes)
	for p, v := range result {
		if reverse[p] != v {
			t.Errorf("Runes were not reversed! expected %s, got %s", reverse, result)
		}
	}
}
