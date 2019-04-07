package chapter4

import "strings"

const Values = "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."

func StringsContains(subStrings string) bool {
	return strings.Contains(Values, subStrings)
}
func StringsCompare(values string, subString string) int {
	return strings.Compare(values, subString)
}
func StringsToUpper(subStrings string) string {
	return strings.ToUpper(subStrings)
}

func StringsToLower(subStrings string) string {
	return strings.ToLower(subStrings)
}

func StringsToTitle(subStrings string) string {
	return strings.ToTitle(subStrings)
}

var UpperCase = map[string]string{
	"a": "A",
	"b": "B",
}

var LowerCase = map[string]string{
	"A": "a",
	"B": "b",
}

func StringsCount(subStrings string) int {
	return strings.Count(Values, subStrings)
}
func StringsSplit(split string) []string {
	return strings.Split(Values, split)
}

func StringsJoin(subStrings []string) string {
	return strings.Join(subStrings, " ")
}

func StringsIndex(subStrings string) int {
	return strings.Index(Values, subStrings)
}

func StringsHasPrefix(subStrings string) bool {
	return strings.HasPrefix(Values, subStrings)
}

func StringsHasSuffix(subStrings string) bool {
	return strings.HasSuffix(Values, subStrings)
}
func StringsTrim(values string) string {
	return strings.TrimSpace(values)
}

func StringsReplacer(values string) string {
	newReplacer := strings.NewReplacer("\n", "", "\t", "", " ", "")
	return newReplacer.Replace(values)
}
