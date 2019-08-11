package assistance

import (
	"strconv"
	"strings"
)

func ToInt(value string) int {
	if value == "" {
		return 0
	}
	v := strings.TrimSpace(value)
	i, _ := strconv.Atoi(v)
	return i
}
