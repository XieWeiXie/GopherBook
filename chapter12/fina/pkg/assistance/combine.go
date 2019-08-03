package assistance

import (
	"GopherBook/chapter12/fina/configs"
	"fmt"
)

func FullURL(val string) string {
	root := configs.RootURL[:len(configs.RootURL)-5]
	return fmt.Sprintf("%s%s", root, val)
}
