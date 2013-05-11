package matcher

import "testing"

// build the matcher required for testing
func BuildMatcher() Matcher {
	strings := []string {
		"adult",
		"fear",
		"here",
		"kitten",
		"sitting",
		"test",
		"yusuke",
	}
	return NewMatcher(strings)
}

// test's the matchers Closest method
func TestClosest(t *testing.T) {
	m := BuildMatcher()

	if m.Closest("ktn") != "kitten" {
		t.Errorf("ktn should match 'kitten'! it was %s instead.", m.Closest("ktn"))
	}
	if m.Closest("t") != "test" {
		t.Errorf("ktn should match 'test'! it was %s instead.", m.Closest("t"))
	}
}

func TestClosestList(t *testing.T) {
	m:= BuildMatcher()

	matcherResult := m.ClosestList("ktn", 4)
	if len(matcherResult) != 4 {
		t.Errorf("matcherResult returned result of incorrect length! expected 4, was %d instead", len(matcherResult))
	}
}

func TestClosestListMax(t *testing.T) {
	m:= BuildMatcher()

	matcherResult := m.ClosestList("ad", m.Length + 1)
	if len(matcherResult) != m.Length {
		t.Errorf("matcherResult returned incorrect result length %d, expected %d", 
			len(matcherResult), m.Length)
	} else if matcherResult[0].Value != "adult" {
		t.Errorf("Closest with length returned incorrect result: %s instead of %s",
			matcherResult[0], "adult")
	}
}

func BenchmarkClosest(t* testing.T) {
	m:= BuildMatcher()
	matcherResult := m.ClosestList("t", 4)
	matcherResult = m.ClosestList("te", 4)
	matcherResult = m.ClosestList("tes", 4)
	matcherResult = m.ClosestList("test", 4)
}
