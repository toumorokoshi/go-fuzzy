package matcher

import "testing"

// build the matcher required for testing
func BuildMatcher() Matcher {
	strings := [...]string {
		"adult",
		"fear",
		"here",
		"kitten",
		"sitting",
		"test",
		"yusuke",
	}
	return Matcher{strings}
}

// test's the matchers Closest method
func TestClosest(t *testing.T) {
	matcher := BuildMatcher()

	if matcher.Closest("ktn") != "kitten" {
		t.Errorf("ktn should match 'kitten'! it was %s instead.", matcher.Closest("ktn"))
	}
	if matcher.Closest("t") != "test" {
		t.Errorf("ktn should match 'test'! it was %s instead.", matcher.Closest("t"))
	}
}

func TestClosestList(t *testing.T) {
	matcher := BuildMatcher()

	matcherResult := matcher.Closest("ktn", 4)
	if len(matcherResult) != 4 {
		t.Errorf("matcherResult returned result of incorrect length! expected 4, was %d instead", len(matcherResult))
	}

	matcherResult := matcher.Closest("ad", matcherResult.Length + 1)
	if len(matcherResult) != matcherResult.Length {
		t.Errorf("matcherResult returned incorrect result length %d, expected %d", 
			len(matcherResult), matcherResult.Length)
	} else if matcherResult[0] != "adult" {
		t.Errorf("Closest with length returned incorrect result: %s instead of %s",
			matcherResult[0], "adult")
	}
}

func BenchmarkClosest(t* testing.T) {
	matcher := BuildMatcher()
	matcherResult := matcher.Closest("t", 4)
	matcherResult := matcher.Closest("te", 4)
	matcherResult := matcher.Closest("tes", 4)
	matcherResult := matcher.Closest("test", 4)
}
