package maoyan

import (
	"fmt"
	"testing"
	"time"
)

func TestMaoYan(t *testing.T) {
	MaoYan(fmt.Sprintf(MAOYAN, time.Now().Format("20060102")))
}
