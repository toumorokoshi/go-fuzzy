/* Provides a matcher object, used to efficently find matches to a
/* query
*/
package fuzzy

import "fmt"
import "sort"

type Matcher struct {
	elements []string
	Length   int
}

type Match struct {
	Value          string
	Levenshtein    int
}

func (m Match) String() string {
	return fmt.Sprintf("{ Levenshtein: %d, Value: %s }", m.Levenshtein, m.Value)
}

type Matches []*Match

// provides methods to make matches sortable
func (m Matches) Len() int           { return len(m) }
func (m Matches) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Matches) Less(i, j int) bool { return m[i].Levenshtein < m[j].Levenshtein }

func NewMatcher(elements []string) Matcher {
	return Matcher{elements, len(elements)}
}

// finds the closest match and returns it
func (m *Matcher) Closest(matchString string) string {
	return m.ClosestList(matchString, 1)[0].Value
}

/* finds the n closest matches and returns them
   start by filtering the requirements:
       * the order of characters appearing in the substring matches the characters in the substring
       * utilize the levenshtein distance after that
*/
func (m *Matcher) ClosestList(matchString string, count int) Matches {
	if count > m.Length {
		count = m.Length
	}
	matchElements := make(Matches, m.Length)
	for pos, element := range m.elements {
		matchElements[pos] = &Match{element, Levenshtein(element, matchString)}
	}
	sort.Sort(matchElements)
	return matchElements[0:count]
}
