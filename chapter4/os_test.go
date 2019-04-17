package chapter4

import "testing"

func TestOS(test *testing.T) {
	OsUsage()
	//OSUsageWith()
	OSPathUsage()
	OSPathWindows()
	OSDirUsage()
	OSExecUsage()
}
