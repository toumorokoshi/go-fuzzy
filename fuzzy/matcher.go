/* Provides a matcher object, used to efficently find matches to a
/* query
*/
package fuzzy

import "sort"

type Matcher struct {
	elements []string
	Length   int
}

type Match struct {
	Value string
	Score   int
}

type Matches []*Match

// provides methods to make matches sortable
func (m Matches) Len() int           { return len(m) }
func (m Matches) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Matches) Less(i, j int) bool { return m[i].Score > m[j].Score }

func NewMatcher(elements []string) Matcher {
	return Matcher{elements, len(elements)}
}

// finds the closest match and returns it
func (m *Matcher) Closest(matchString string) string {
	return m.ClosestList(matchString, 1)[0].Value
}

/* finds the n closest matches and returns them
 */
func (m *Matcher) ClosestList(matchString string, count int) Matches{
	matchElements := make(Matches, m.Length)
	for pos, element := range m.elements {
		matchElements[pos] = &Match{element, Levenshtein(element, matchString)}
	}
	sort.Sort(matchElements)
	return matchElements[0:count]
}
