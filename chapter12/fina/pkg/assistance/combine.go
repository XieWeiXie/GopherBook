package assistance

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/configs"
)

func FullURL(val string) string {
	root := configs.RootURL[:len(configs.RootURL)-5]
	return fmt.Sprintf("%s%s", root, val)
}
