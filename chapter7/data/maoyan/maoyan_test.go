package maoyan

import (
	"fmt"
	"testing"
	"time"
)

func TestMaoYan(t *testing.T) {
	// DEPRECATED: 猫眼接口变更: 2019-12-05
	// MaoYan(fmt.Sprintf(MAOYAN, time.Now().Format("20060102")))
	MaoYan(fmt.Sprintf(MAOYANNEW, time.Now().Format("20060102")))
}
