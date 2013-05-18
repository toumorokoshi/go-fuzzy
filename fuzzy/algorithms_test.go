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
		TestStringSubsequence{"dDe", "disasterDeafening", true},
	}
	for _, v := range testStrings {
		matched := SequenceMatchCaseSensitive(v.A, v.B)
		if matched != v.Expected {
			t.Errorf("Levenshtein distance between %s and %s should be %t, was %t", v.A, v.B, v.Expected, matched)
		}
	}
}

func TestOrderSignificance(t *testing.T) {
	testStrings := [...]TestString{
		TestString{"rdn", "relaxeddailyroutine", 3},
		TestString{"rdn", "RelaxedDailyroutine", 5},
		TestString{"rdn", "relaxedDailyRoutiNe", 5},
		TestString{"npe", "electroplankton", 0},
		TestString{"mastermind", "masterminder", 10},
		TestString{"dDe", "disasterdeafening", 3},
		TestString{"t", "adult", 1},
		TestString{"dDe", "disasterDeafening", 4},
	}
	for _, v := range testStrings {
		t.Logf("Testing OrderSignificance %s against %s...", v.A, v.B)
		value := OrderSignificance(v.A, v.B)
		if value != v.Expected {
			t.Errorf("Order and Significance value between %s and %s should be %d, was %d", v.A, v.B, v.Expected, value)
		}
	}
}
