package assistance

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

/**/
func words(value string) (string, string) {
	var buf bytes.Buffer
	var bufNotHan bytes.Buffer
	for _, i := range value {
		if unicode.Is(unicode.Scripts["Han"], i) {
			buf.WriteRune(i)
		} else {
			bufNotHan.WriteRune(i)
		}
	}
	return buf.String(), bufNotHan.String()
}

func GetWordsExceptHan(value string) string {
	_, notHan := words(value)
	return notHan
}

func GetWordsHan(value string) string {
	han, _ := words(value)
	return han

}

/**/
func GetDate(value string) (time.Time, time.Time, error) {
	// 2019年07月12日-07月28日(17天)
	list := strings.SplitN(value, "-", -1)
	if len(list) < 1 {
		return time.Time{}, time.Time{}, fmt.Errorf("error")
	}
	reg, err := regexp.Compile(`\d+`)
	listString := reg.FindAllString(list[0], -1)
	if err != nil || len(listString) != 3 {
		return time.Time{}, time.Time{}, err
	}
	year, month, day := listString[0], listString[1], listString[2]
	start, _ := time.ParseInLocation("2006-01-02", fmt.Sprintf("%s-%s-%s", year, month, day), time.Local)
	end := start.AddDate(0, 0, 16)
	return start, end, nil
}

/**/

func GetDisciplines(value string) []string {
	return strings.Split(value, "、")
}

/**/

func ToInt(value string) int {
	result, _ := strconv.Atoi(value)
	return result
}

/**/

func ReplaceSpace(value string) string {
	replacer := strings.NewReplacer(" ", "", "\n", "")
	return replacer.Replace(value)
}

/**/

func SplitBYColon(value string, sep string) string {
	list := strings.Split(value, sep)
	if len(list) != 2 {
		return "-1"
	}
	return list[1]
}

/**/

func SplitBySep(value string, sep string) []string {
	replacer := strings.NewReplacer("\r", "", "\n", "")
	return strings.Split(replacer.Replace(value), sep)
}

/**/
