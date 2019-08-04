package configs

import "fmt"

// URL Collections
var (
	RootURL = "https://www.fina-gwangju2019.com/chn/"
)
var (
	MatchDescription = fmt.Sprintf("%s%s", RootURL, "contentsView.do?pageId=chn3")
	MatchSymbol      = fmt.Sprintf("%s%s", RootURL, "contentsView.do?pageId=chn4")
	MatchProjects    = fmt.Sprintf("%s%s", RootURL, "contentsView.do?pageId=chn29")
	MatchMessage     = fmt.Sprintf("%s%s", RootURL, "contentsView.do?pageId=chn91")
	MatchHistory     = fmt.Sprintf("%s%s", RootURL, "contentsView.do?pageId=chn15")
	MatchBrief       = fmt.Sprintf("%s%s", RootURL, "contentsView.do?pageId=chn14")
)

var (
	MatchSports    = "https://www.fina-gwangju2019.com/chn/contentsView.do?pageId=chn%d&sn=%d"
	MatchSportsMap = map[int]int{}
	MatchPostDo    = "https://www.fina-gwangju2019.com/pg/sportEntriesData.do"
)

func init() {
	MatchSportsMap = make(map[int]int)
	for i := 0; i < 7; i++ {
		MatchSportsMap[36-i] = i
	}
}
