package assistance

import "testing"

func TestDownloadByChromeHeadless(t *testing.T) {
	DownloadByChromeHeadless("https://www.baidu.com")
}
