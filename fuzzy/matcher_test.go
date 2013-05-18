package fuzzy

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
		"ReadingRainbow",
    "EEResume.pdf",
		"Lab5-ASSUBMITTED.docx",
		"mortgage-calculator-ui-mobile.js",
		"mobile.js",
	}
	return NewMatcher(strings)
}

type TestClosestStruct struct {
	TestString, Expected string
}

func TestComparisons(t *testing.T) {
	m1 := NewMatch("ReadingRainbow", "rr")
	m2 := NewMatch("fear", "rr")
	matches := Matches{m1, m2}
	if matches.Less(1, 0) {
		t.Errorf("ReadingRainbow should be closer than fear to rr!")
	}
}

// test's the matchers Closest method
func TestClosest(t *testing.T) {
	m := BuildMatcher()
	testClosestStructs := [...]TestClosestStruct{
		TestClosestStruct{"ktn", "kitten"},
		TestClosestStruct{"tes", "test"},
		TestClosestStruct{"rr", "ReadingRainbow"},
		TestClosestStruct{"resume", "EEResume.pdf"},
		TestClosestStruct{"mobile.js", "mobile.js"},
	}
	for _, v := range testClosestStructs {
		t.Logf("Testing match %s...", v.TestString)
		if m.Closest(v.TestString) != v.Expected {
			t.Errorf("%s should match '%s'! it was %s instead.", v.TestString, v.Expected, m.Closest(v.TestString))
		}
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

	matcherResult := m.ClosestList("adu", m.Length + 1)
	if len(matcherResult) != m.Length {
		t.Errorf("matcherResult returned incorrect result length %d, expected %d", 
			len(matcherResult), m.Length)
	} else if matcherResult[0].Value != "adult" {
		t.Errorf("Closest with length returned incorrect result: %s instead of %s",
			matcherResult[0], "adult")
	}
}

func BenchmarkClosest(b* testing.B) {
	m:= BuildMatcher()
	for i :=0; i <= b.N; i++ {
		_ = m.ClosestList("t", 4)
		_ = m.ClosestList("te", 4)
		_ = m.ClosestList("tes", 4)
		_ = m.ClosestList("test", 4)
	}
}
