package model

import (
	"time"
)

type Base struct {
	Id        int64      `xorm:"pk notnull"`
	CreatedAt time.Time  `xorm:"created" json:"created_at"`
	UpdatedAt time.Time  `xorm:"updated" json:"updated_at"`
	DeletedAt *time.Time `xorm:"deleted" json:"deleted_at"`
}

type VoteByXORM struct {
	Base        `xorm:"extends"`
	Title       string `xorm:"varchar(10) notnull" json:"title"`
	AdminId     uint   `xorm:"index" json:"admin_id"`
	Description string `xorm:"varchar(64) default(null)" json:"description"`
	Choice      []Choice
	DeadLine    time.Time `xorm:"timestamp" json:"dead_line"`
	IsAnonymous bool
	IsSingle    bool
}

func (v VoteByXORM) TableName() string {
	return "vote_by_xorm"
}

type ChoiceByXORM struct {
	Base   `xorm:"extends"`
	VoteId uint
	Title  string `xorm:"varchar(10) notnull" json:"title"`
	Number int    `xorm:"int(4)" json:"number"`
}

func (v ChoiceByXORM) TableName() string {
	return "choice_by_xorm"
}
