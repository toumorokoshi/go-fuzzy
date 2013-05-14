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

type TestStringSubsequence struct {
	A, B     string
	Expected bool
}

func TestSubsequenceMatch(t *testing.T) {
	testStrings := [...]TestStringSubsequence{
		TestStringSubsequence{"rdn", "relaxeddailyroutine", true},
		TestStringSubsequence{"npe", "electroplankton", false},
		TestStringSubsequence{"mastermind", "masterminder", true},
		TestStringSubsequence{"dDe", "disasterdeafening", false},
		TestStringSubsequence{"dDe", "disasterDeafening", true},
	}
	for _, v := range testStrings {
		matched := SequenceMatchCaseSensitive(v.A, v.B)
		if matched != v.Expected {
			t.Errorf("Levenshtein distance between %s and %s should be %t, was %t", v.A, v.B, v.Expected, matched)
		}
	}
}
