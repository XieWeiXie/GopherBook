package model_v1

import (
	"fmt"
	"strings"
)

const (
	V0 = iota
	V1
	V2
	V3
	V4
)

var validity = make(map[int]struct {
	Time    int
	Comment int
})

func init() {
	validity[V0] = struct {
		Time    int
		Comment int
	}{Time: 0, Comment: 0}

	validity[V1] = struct {
		Time    int
		Comment int
	}{Time: 1, Comment: 30}

	validity[V2] = struct {
		Time    int
		Comment int
	}{Time: 1, Comment: 169}

	validity[V3] = struct {
		Time    int
		Comment int
	}{Time: 1, Comment: 300}

	validity[V4] = struct {
		Time    int
		Comment int
	}{Time: 1, Comment: 500}
}

type VipMember struct {
	base      `xorm:"extends"`
	LevelName string `xorm:"varchar(2) notnull unique 'level_name'" json:"level_name"`
	Start     int    `json:"start"`
	End       int    `json:"end"`
	Points    float64
	Comment   string `xorm:"varchar(128) notnull" json:"comment"`
	Period    int    `json:"period"`
	ToValue   int    `json:"to_value"`
}

func (VipMember) TableName() string {
	return "beeQuick_vip_member"
}

type VipMemberSerializer struct {
	ID        uint   `json:"id"`
	LevelName string `json:"level_name"`
	Start     int    `json:"start"`
	End       int    `json:"end"`
	Comment   string `json:"comment"`
	Period    int    `json:"period"`
	ToValue   int    `json:"to_value"`
}

func (vip VipMember) Serializer() VipMemberSerializer {
	return VipMemberSerializer{
		ID:        vip.ID,
		LevelName: vip.LevelName,
		Start:     vip.Start,
		End:       vip.End,
		Comment:   vip.Comment,
		Period:    vip.Period,
		ToValue:   vip.ToValue,
	}
}

func DefaultVipMemberRecord() []*VipMember {
	return []*VipMember{
		{
			LevelName: strings.ToUpper("v0"),
			Start:     0,
			End:       29,
			Points:    0.5,
			Comment:   fmt.Sprintf("获取%.2f倍积分", 0.5),
			Period:    0,
			ToValue:   0,
		},
		{
			LevelName: strings.ToUpper("v1"),
			Start:     30,
			End:       198,
			Points:    1.0,
			Comment:   fmt.Sprintf("获取%.2f倍积分", 1.0),
			Period:    1,
			ToValue:   30,
		},
		{
			LevelName: strings.ToUpper("v2"),
			Start:     199,
			End:       498,
			Points:    1.5,
			Comment:   fmt.Sprintf("获取%.2f倍积分", 1.5),
			Period:    1,
			ToValue:   169,
		},
		{
			LevelName: strings.ToUpper("v3"),
			Start:     499,
			End:       998,
			Points:    2.0,
			Comment:   fmt.Sprintf("获取%.2f倍积分", 2.0),
			Period:    1,
			ToValue:   300,
		},
		{
			LevelName: strings.ToUpper("v4"),
			Start:     999,
			End:       0,
			Points:    3.0,
			Comment:   fmt.Sprintf("获取%.2f倍积分", 3.0),
			Period:    1,
			ToValue:   500,
		},
	}
}
