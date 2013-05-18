/* 
Provides a matcher object, used to efficently find matches to a query

The matcher designed to be flexible, but it is being tested extensively against file searching. In order of importance:
* subsequence match after 6 characters
* Matches against 'significant' characters, such as capitals or characters after a - or _
* Filename matches. e.g. matches against the actual name of the file

*/
package fuzzy

import "bytes"
import "fmt"
import "sort"
import "strings"

type Matcher struct {
	elements []string
	Length   int
}

type Match struct {
	Value       string
	Exact bool
	SignificantMatch bool
	IndexMatch int
	OrderSignificance int
	Levenshtein int
}

func (m Match) String() string {
	return fmt.Sprintf("{ Value: %s, Exact: %t, SignificantMatch: %t, IndexMatch: %d, OrderSignificance: %d, Levenshtein: %d }", 
		m.Value, m.Exact, m.SignificantMatch, m.IndexMatch, m.OrderSignificance, m.Levenshtein)
}

type Matches []*Match


func (m Matches) String() string {
	var buffer bytes.Buffer
	for _, v := range m {
		buffer.WriteString(v.String())
		buffer.WriteString("\n")
	}
	return buffer.String()
}

// provides methods to make matches sortable
func (m Matches) Len() int           { return len(m) }
func (m Matches) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Matches) Less(i, j int) bool { 
	if m[i].Exact != m[j].Exact { return false }
	if m[i].SignificantMatch != m[j].SignificantMatch { return m[i].SignificantMatch }
	if m[i].IndexMatch != m[j].IndexMatch { 
		return m[j].IndexMatch == -1 ||
			(m[i].IndexMatch != -1 && m[i].IndexMatch < m[j].IndexMatch)
	}
	if m[i].OrderSignificance != m[j].OrderSignificance {
	  return m[i].OrderSignificance > m[j].OrderSignificance
  }
	return m[i].Levenshtein < m[j].Levenshtein
}

func NewMatcher(elements []string) Matcher {
	return Matcher{elements, len(elements)}
}

// finds the closest match and returns it
func (m *Matcher) Closest(matchString string) string {
	return m.ClosestList(matchString, 1)[0].Value
}

func NewMatch(element, matchString string) *Match {
	indexMatch := strings.Index(element, matchString)
	significantMatch := OrderSignificance(element, matchString)
	return &Match{element, 
		indexMatch == 0 && len(element) == len(matchString),
		significantMatch == len(matchString) * 2,
		indexMatch,
		significantMatch,
		Levenshtein(element, matchString)}
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
		matchElements[pos] = NewMatch(element, matchString)
	}
	sort.Sort(matchElements)
	return matchElements[0:count]
}
