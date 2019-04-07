package chapter4

import (
	"fmt"
	"time"
)

func TimeUsage() {
	now := time.Now()

	// 获取年
	fmt.Println(now.Year())
	// 获取月份
	fmt.Println(now.Month())
	// 获取日前
	fmt.Println(now.Date())
	// 获取天
	fmt.Println(now.Day())
	// 小时
	fmt.Println(now.Hour())
	// 分
	fmt.Println(now.Minute())
	// 秒
	fmt.Println(now.Second())
	// 毫秒
	fmt.Println(now.Unix())
	// 纳秒
	fmt.Println(now.UnixNano())
}

func TimeOperate() {
	start := time.Now()
	time.Sleep(1 * time.Second)
	// 两个时间差
	fmt.Println(time.Now().Sub(start))

	// 格式化
	fmt.Println(start.Format("2006-01-02 15:04:06"))

	// 截取
	fmt.Println(start.Round(time.Second))
	fmt.Println(start.Truncate(time.Second))

	stringTime := "1991-12-25 19:00:00"
	birthday, _ := time.ParseInLocation("2006-01-02 15:04:05", stringTime, time.Local)
	fmt.Println(birthday.String())
}

func TimeAdd() {
	now := time.Now()

	// 一天前
	oneDayBefore := now.AddDate(0, 0, -1)

	fmt.Println(now.String(), oneDayBefore.String())

	// 一小时前
	oneHourBefore := now.Add(-1 * time.Hour)
	fmt.Println(oneHourBefore)
}
