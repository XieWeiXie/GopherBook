package bilibili

type ResultForBiliBili struct {
	Title     string `json:"title"`
	Href      string `json:"href"`
	Author    string `json:"author"`
	AuthorURL string `json:"author_url"`
	Play      string `json:"play"`
	View      string `json:"view"`
	Pts       string `json:"pts"`
}
