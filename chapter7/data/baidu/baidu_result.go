package baidu

import "github.com/jinzhu/gorm"

type ResultBaiDu struct {
	gorm.Model
	Keyword string `json:"keyword" gorm:"type:varchar(32)"`
	Href    string `json:"href" gorm:"type:varchar(256)"`
	Number  int    `json:"number" gorm:"type:integer(11)"`
}

func (R ResultBaiDu) TableName() string {
	return "result_baidu"
}
