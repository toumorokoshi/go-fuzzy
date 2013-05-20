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
	return NewMatcherFromStrings(strings)
}

func BuildMatcherWithData() Matcher {
	matchStructs := []MatchStruct {
		MatchStruct{"adult", nil},
		MatchStruct{"Fear", map[string]string{"path": "test"}},
		MatchStruct{"Here", map[string]string{"path": "test"}},
		MatchStruct{"Kitten", map[string]string{"path": "test"}},
		MatchStruct{"sitting", map[string]string{"path": "test"}},
		MatchStruct{"test", map[string]string{"path": "test"}},
		MatchStruct{"yusuke", map[string]string{"path": "test"}},
		MatchStruct{"ReadingRainbow", map[string]string{"path": "test"}},
		MatchStruct{"EEResume.pdf", map[string]string{"path": "test"}},
		MatchStruct{"Lab5-ASSUBMITTED.docx", map[string]string{"path": "test"}},
		MatchStruct{"mortgage-calculator-ui-mobile.js", map[string]string{"path": "test"}},
		MatchStruct{"mobile.js", map[string]string{"path": "test"}},
	}
	return NewMatcher(matchStructs)
}

// test the build matcher, with data
func TestBuildMatcherWithData (t *testing.T) {
	m := BuildMatcherWithData()
	if m.Closest("adult").Data != nil {
		t.Errorf("Data for adult should be null!")
	}
	if m.Closest("Fear").Data["path"] != "test" {
		t.Errorf("Data for Fear should have an element path with value test!")
	}
}

type TestClosestStruct struct {
	TestString, Expected string
}

func TestComparisons(t *testing.T) {
	m1 := NewMatchFromString("ReadingRainbow", "rr")
	m2 := NewMatchFromString("fear", "rr")
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
		if m.Closest(v.TestString).Value != v.Expected {
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
