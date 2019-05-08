package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

type Person struct {
	Avatar         string    `json:"avatar"`
	OriginWeChatID string    `json:"origin_id"`
	ID             uint      `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	Telephone      string    `json:"telephone"`
	Gender         int       `json:"gender"` // 男1 女0
	WhatIsUp       string    `json:"what's_up"`
}

func uuid() string {
	b := []byte(fmt.Sprintf("%v", time.Now().Unix()))
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func NewPersonRecords() []Person {
	var persons []Person
	persons = []Person{
		{
			Avatar:         "http://images.org/123",
			OriginWeChatID: uuid(),
			ID:             1,
			CreatedAt:      time.Now(),
			Telephone:      "1234567890",
			Gender:         1,
			WhatIsUp:       "Hello Golang",
		},
		{
			Avatar:         "http://images.org/123",
			OriginWeChatID: uuid(),
			ID:             2,
			CreatedAt:      time.Now(),
			Telephone:      "987654321",
			Gender:         0,
			WhatIsUp:       "Hello Python",
		},
	}
	return persons
}
