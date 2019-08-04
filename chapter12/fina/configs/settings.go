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

var (
	MatchRank = "https://www.fina-gwangju2019.com/pg/getEventRanking.do"
)

var (
	MatchHistoryYear = "https://sportapi.widgets.sports.gracenote.com/games_v2/getseasonlist/competitionsetid/10/languagecode/2.json"
	MatchPostEvent   = "https://sportapi.widgets.sports.gracenote.com/games_v2/getmedaltable_season/competitionsetid/10/season/%d/languagecode/2.json"
)

var (
	MatchRecords = "https://sportapi.widgets.sports.gracenote.com/games_v2/getrecordlist/sportid/117/recordtype/world/competitionsetid/10/languagecode/2.json"
)
