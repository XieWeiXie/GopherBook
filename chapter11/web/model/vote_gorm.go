package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Vote struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(32)"`
	AdminId     uint   `json:"admin_id"`
	Description string `json:"description" gorm:"type:varchar(64)"`
	Choice      []Choice
	DeadLine    time.Time
	IsAnonymous bool
	IsSingle    bool
}

type VoteSerializer struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	AdminId     uint      `json:"admin_id"`
	Description string    `json:"description"`
	DeadLine    string    `json:"dead_line"`
	IsAnonymous string    `json:"is_anonymous"`
	IsSingle    string    `json:"is_single"`
}

func (v Vote) Serializer() VoteSerializer {

	var isAnonymous = func(key bool) string {
		if key {
			return "匿名"
		}
		return "公开"

	}
	var isSingle = func(key bool) string {
		if key {
			return "单项选择"
		}
		return "多项选择"
	}

	return VoteSerializer{
		ID:          v.ID,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		Title:       v.Title,
		AdminId:     v.AdminId,
		Description: v.Description,
		DeadLine:    v.DeadLine.Format("2006-01-02 15:04:05"),
		IsAnonymous: isAnonymous(v.IsAnonymous),
		IsSingle:    isSingle(v.IsSingle),
	}
}

type Choice struct {
	gorm.Model
	VoteId uint
	Title  string `gorm:"type:varchar(32)" json:"title"`
	Number int    `gorm:"type:integer(4)" json:"number"`
}
