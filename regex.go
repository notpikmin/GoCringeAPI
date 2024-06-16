package CringeApi

import (
	"github.com/dlclark/regexp2"
)

func Regexp2FindAllString(re *regexp2.Regexp, s string) []*regexp2.Match {
	var matches []*regexp2.Match
	m, _ := re.FindStringMatch(s)
	if m != nil {

		//fmt.Println(m.Capture.String())
	}
	for m != nil {
		matches = append(matches, m)
		m, _ = re.FindNextMatch(m)

		if m != nil {

			//	fmt.Println(m.Capture.String())
		}
	}
	return matches
}

func MatchAllWords(bio string) []*regexp2.Match {
	out := make([]*regexp2.Match, 0)
	for _, item := range CringeMetric {
		out = append(out, Regexp2FindAllString(item.Pattern, bio)...)
	}
	return out
}
