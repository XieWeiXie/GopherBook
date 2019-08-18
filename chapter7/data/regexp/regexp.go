package regexpnode

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type Element struct {
	TagName     string            `json:"tag_name"`
	TextContent string            `json:"text_content"`
	Attrs       map[string]string `json:"attrs"`
}

var element = `<p(.*?)>(.*?)</p>`
var attr = `(.*?)="(.*?)"`

func ParseByRegexp() {

	content, err := ioutil.ReadFile("node.html")
	if err != nil {
		log.Println(err)
		return
	}
	re, err := regexp.Compile(element)
	if err != nil {
		log.Println(err)
		return
	}
	results := re.FindAllStringSubmatch(string(content), -1)
	for _, i := range results {
		var one Element
		one.TagName = "p"
		one.TextContent = i[len(i)-1]
		if len(i) > 2 {
			var attrs = make(map[string]string)
			attrReg := regexp.MustCompile(attr)
			for k := 1; k < len(i); k++ {
				a := attrReg.FindAllStringSubmatch(i[k], -1)
				if len(a) != 0 {
					attrs[strings.TrimSpace(a[0][1])] =
						strings.TrimSpace(a[0][2])
				}
			}
			one.Attrs = attrs
		}
		fmt.Println(one)
	}

}
