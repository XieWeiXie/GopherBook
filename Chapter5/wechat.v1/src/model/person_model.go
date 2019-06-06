package model

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	NickName         string `gorm:"type:varchar" json:"nick_name"`
	CommentName      string `gorm:"type:varchar" json:"comment_name"`
	WeChatAccount    string `gorm:"type:varchar" json:"we_chat_account"`
	Signature        string `gorm:"type:varchar(34)" json:"signature"`
	FromID           uint
	Moments          []Moment
	OfficialAccounts []OfficialAccount
	Addresses        []ReceiveAddress
}
