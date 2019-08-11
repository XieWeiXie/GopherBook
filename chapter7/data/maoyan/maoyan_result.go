package maoyan

type ResultForMaoYan struct {
	AvgSeatView  string `json:"avg_seat_view"`
	AvgShowView  string `json:"avg_show_view"`
	BoxRate      string `json:"box_rate"`
	MovieName    string `json:"movie_name"`
	ReleaseInfo  string `json:"release_info"`
	BoxInfo      string `json:"box_info"`
	ShowInfo     string `json:"show_info"`
	ShowRate     string `json:"show_rate"`
	SplitBoxRate string `json:"split_box_rate"`
	SumBoxInfo   string `json:"sum_box_info"`
}
