package assistance

import (
	"fmt"
	"log"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	PORT = 4444
)

func Selenium(url string) (string, error) {
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
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7",
		},
	}
	caps.AddChrome(chromeCaps)

	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", PORT))
	if err != nil {
		panic(err)
	}
	defer webDriver.Close()

	err = webDriver.Get(url)
	if err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}
	title, _ := webDriver.Title()
	log.Println(title)
	iframe, err := webDriver.FindElement(selenium.ByCSSSelector, "#g_iframe")
	if err != nil {
		panic(err)
	}
	if err := webDriver.SwitchFrame(iframe); err != nil {
		panic(err)
	}
	content, err := webDriver.PageSource()
	if err != nil {
		panic(err)
	}
	return content, nil

}

func SeleniumGetContent(url string) (string, error) {
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", PORT))
	if err != nil {
		panic(err)
	}
	if err := webDriver.Get(url); err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}
	return webDriver.PageSource()
}
