package assistance

import (
	"fmt"
	"io"
	"strings"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func DownloadByChromeHeadless(url string) (io.Reader, error) {
	<-rateTime
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}

	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			"--headless",
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7",
		},
	}
	caps.AddChrome(chromeCaps)

	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 4444))
	if err != nil {
		panic(err)
	}
	defer webDriver.Close()
	err = webDriver.Get(url)
	if err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}
	source, _ := webDriver.PageSource()
	return strings.NewReader(source), nil

}
