package douban

type ResultForDouBan struct {
	Casts     []string `json:"casts"`
	Directors []string `json:"directors"`
	Rate      string   `json:"rate"`
	Star      string   `json:"star"`
	Title     string   `json:"title"`
	Url       string   `json:"url"`
}
