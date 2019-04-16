package chapter4

import "testing"

func TestSimpleUsageForLog(t *testing.T) {
	DefaultUsageForLog()
	SpecialUsageLog()
	SpecialUsageWithBytes()
	SpecialUsageWithFile()
}
