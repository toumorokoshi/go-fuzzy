package fuzzy

import "testing"

type TestString struct {
	A, B     string
	Expected int
}

func TestLevenshtein(t *testing.T) {
	testStrings := [...]TestString{
		TestString{"kitten", "sitting", 3},
		TestString{"same", "same", 0},
	}
	for _, v := range testStrings {
		distance := Levenshtein(v.A, v.B)
		if distance != v.Expected {
			t.Errorf("Levenshtein distance between %s and %s should be %d, was %d", v.A, v.B, v.Expected, distance)
		}
	}
}
