package models

import "time"

type Base struct {
	Id        int64      `xorm:"notnull pk autoincr 'id'" json:"id"`
	CreatedAt time.Time  `xorm:"created" json:"created_at"`
	UpdatedAt time.Time  `xorm:"updated" json:"updated_at"`
	DeletedAt *time.Time `xorm:"deleted" json:"deleted_at"`
}
