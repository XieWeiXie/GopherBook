package wangyiyun

type ResultForWangYiYun struct {
	Title          string `json:"title"`
	Url            string `json:"url"`
	UpdateFrequent string `json:"update_frequent"`
	LastUpdate     string `json:"last_update"`
	Ranks          []Rank `json:"ranks"`
	Fav            int    `json:"fav"`
	Share          int    `json:"share"`
	Comment        int    `json:"comment"`
}

type Rank struct {
	Title     string `json:"title"`
	Url       string `json:"url"`
	Time      string `json:"time"`
	Author    string `json:"author"`
	AuthorUrl string `json:"author_url"`
}
