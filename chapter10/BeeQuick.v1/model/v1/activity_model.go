package model_v1

import "time"

type Activity struct {
	base    `xorm:"extends"`
	Name    string    `xorm:"varchar(32)" json:"name"`
	Title   string    `xorm:"varchar(32)" json:"title"`
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
	Avatar  string    `xorm:"varchar(255)" json:"avatar"`
	ShopIds []int     `xorm:"blob" json:"shop_ids"`
	Status  int       `xorm:"varchar(10)"`
}

func (a Activity) TableName() string {
	return "beeQuick_activity"
}

type ActivitySerializer struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	Avatar    string    `json:"avatar"`
	ShopIds   []int     `json:"shop_ids"`
	Status    string    `json:"status"`
}

func (a Activity) Serializer() ActivitySerializer {
	return ActivitySerializer{
		Id:        a.ID,
		CreatedAt: a.CreatedAt.Truncate(time.Second),
		UpdatedAt: a.UpdatedAt.Truncate(time.Second),
		Name:      a.Name,
		Title:     a.Title,
		Start:     a.Start,
		End:       a.End,
		Avatar:    a.Avatar,
		ShopIds:   a.ShopIds,
		Status:    ActivityStatus[a.Status],
	}
}

const (
	DOING = iota
	PROGRESSING
	CANCEL
	FINISH
	ADVANCE
)

var ActivityStatus = make(map[int]string)
var ActivityStatusEn = make(map[int]string)

func init() {
	ActivityStatus[DOING] = "未开始"
	ActivityStatus[PROGRESSING] = "进行中"
	ActivityStatus[CANCEL] = "取消"
	ActivityStatus[FINISH] = "结束"
	ActivityStatus[ADVANCE] = "提前"

	ActivityStatusEn[DOING] = "DOING"
	ActivityStatusEn[PROGRESSING] = "PROGRESSING"
	ActivityStatusEn[CANCEL] = "CANCEL"
	ActivityStatusEn[FINISH] = "FINISH"
	ActivityStatusEn[ADVANCE] = "ADVANCE"

}

type Activity2Product struct {
	ProductId  int64 `xorm:"index"`
	ActivityId int64 `xorm:"index"`
}

func (s Activity2Product) TableName() string {
	return "beeQuick_activity2Product"
}

type Shop2Activity struct {
	ShopId     int64 `xorm:"index"`
	ActivityId int64 `xorm:"index"`
}

func (s Shop2Activity) TableName() string {
	return "beeQuick_shop2Activity"
}
