package wangyiyun

import (
	"testing"
)

func TestWangYiYun(t *testing.T) {
	//fmt.Println(assistance.Selenium(ROOTURL))
	WangYiYun(ROOTURL)
}

func TestWangYiYunByCdp(t *testing.T) {
	WangYiYunByChromedp("https://music.163.com/#/discover/toplist")
}
