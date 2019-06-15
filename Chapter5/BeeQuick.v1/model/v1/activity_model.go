package model_v1

import "time"

type Activity struct {
	base   `xorm:"extends"`
	Name   string    `xorm:"varchar(32)" json:"name"`
	Title  string    `xorm:"varchar(32)" json:"title"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	Avatar string    `xorm:"varchar(255)" json:"avatar"`
	ShopId int64     `xorm:"index"`
	Shop   Shop      `xorm:"-"`
	Status int       `xorm:"varchar(10)"`
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
	ShopId    int64     `json:"shop_id"`
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
		ShopId:    a.ShopId,
		Status:    activityStatus[a.Status],
	}
}

const (
	DOING = iota
	PROGRESSING
	CANCEL
	FINISH
	ADVANCE
)

var activityStatus = make(map[int]string)

func init() {
	activityStatus[DOING] = "未开始"
	activityStatus[PROGRESSING] = "进行中"
	activityStatus[CANCEL] = "取消"
	activityStatus[FINISH] = "结束"
	activityStatus[ADVANCE] = "提前"
}

type Shop2Activity struct {
	ProductId  int64 `xorm:"index"`
	ActivityId int64 `xorm:"index"`
}
