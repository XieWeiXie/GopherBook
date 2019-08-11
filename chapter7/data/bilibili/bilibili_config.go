package bilibili

import (
	"fmt"
)

const (
	ALL = iota
	ORIGIN
	BANGUMI
	CINEMA
	ROOKIE
)

var (
	HOST    = "https://www.bilibili.com"
	RANKING = fmt.Sprintf(HOST+"%s", "/ranking/")
)
var RankMenu map[int][]string

func init() {
	RankMenu = make(map[int][]string)
	RankMenu[ALL] = []string{"all", "全站榜"}
	RankMenu[ORIGIN] = []string{"origin", "原创榜"}
	RankMenu[BANGUMI] = []string{"bangumi", "新番榜"}
	RankMenu[CINEMA] = []string{"cinema", "影视榜"}
	RankMenu[ROOKIE] = []string{"rookie", "新人榜"}
}

var lists = []struct {
	name string
	tid  int
}{
	{
		name: "全站",
		tid:  0,
	}, {
		name: "动画",
		tid:  1,
	}, {
		name: "国创相关",
		tid:  168,
	}, {
		name: "音乐",
		tid:  3,
	}, {
		name: "舞蹈",
		tid:  129,
	}, {
		name: "游戏",
		tid:  4,
	}, {
		name: "科技",
		tid:  36,
	}, {
		name: "数码",
		tid:  188,
	}, {
		name: "生活",
		tid:  160,
	}, {
		name: "鬼畜",
		tid:  119,
	}, {
		name: "时尚",
		tid:  155,
	}, {
		name: "娱乐",
		tid:  5,
	}, {
		name: "影视",
		tid:  181,
	},
}

var I = []struct {
	name  string
	value int
}{
	{
		name:  "日排行",
		value: 1,
	}, {
		name:  "三日排行",
		value: 3,
	}, {
		name:  "周排行",
		value: 7,
	}, {
		name:  "月排行",
		value: 30,
	},
}

func BiliBiliURL() []string {
	var urlList []string
	for k, v := range RankMenu {
		for _, i := range lists {
			for _, j := range I {
				URL := fmt.Sprintf(RANKING+v[0]+"/%d/%d/%d", i.tid, k, j.value)
				fmt.Println(URL)
				urlList = append(urlList, URL)

			}

		}
	}
	return urlList
}
