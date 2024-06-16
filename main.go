package CringeApi

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"strings"
)

func CheckForErr(err error) bool {
	if err != nil {
		fmt.Println(err)
	}
	return err != nil
}

func ParseBio(b string) (string, int) {
	DownloadCringeMetric()
	matches := MatchAllWords(b)
	score := ScoreBio(b)
	bio := HighlightKeywords(b, matches)
	return bio, score
}
func HighlightKeywords(bio string, matches []*regexp2.Match) string {
	replacer := strings.NewReplacer("*", "\\*", "_", "\\_", "~", "\\~", "`", "\\`", ">", "\\>", "-", "\\-", "#", "\\#")
	bio = replacer.Replace(bio)
	for _, match := range matches {
		bio = strings.ReplaceAll(bio, match.String(), "__**"+match.String()+"**__")
	}
	return bio
}

func ScoreBio(bio string) int {
	score := 0
	for _, item := range CringeMetric {

		score += item.Score * len(Regexp2FindAllString(item.Pattern, bio))
	}
	return score
}
