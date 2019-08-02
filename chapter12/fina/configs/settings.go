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
)
