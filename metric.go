package GoCringeAPI

import (
	"github.com/dlclark/regexp2"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var CringeMetric []CringeItem

type CringeItem struct {
	Score   int
	Pattern *regexp2.Regexp
}

func DownloadCringeMetric() {

	fullURLFile := "https://pastebin.pl/view/raw/c84cab1f"
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	CringeMetric = make([]CringeItem, 0)
	resp, err := client.Get(fullURLFile)
	CheckForErr(err)
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	response := strings.Split(string(b), "\n")
	for _, s := range response {
		ci := strings.Split(s, ":")
		if len(ci) < 2 {
			continue
		}
		score, err := strconv.Atoi(strings.Trim(ci[1], "\r"))
		CheckForErr(err)
		p := ci[0]
		p = strings.ReplaceAll(p, "[lb]", "(?<![^_\\W])")
		p = strings.ReplaceAll(p, "[la]", "(?![^_\\W])")
		p = "(?i)" + p
		pattern, err := regexp2.Compile(p, 0)
		CheckForErr(err)
		cringe := CringeItem{
			Score:   score,
			Pattern: pattern,
		}
		CringeMetric = append(CringeMetric, cringe)
	}
}
