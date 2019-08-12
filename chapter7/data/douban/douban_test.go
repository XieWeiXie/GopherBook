package douban

import (
	"fmt"
	"testing"
)

func TestGetDonBan(t *testing.T) {
	fmt.Println(1)
	GetDouBan(HOST)
	fmt.Println(2)
}

func TestGetDouBanByAPI(t *testing.T) {
	GetDouBanByAPI(API, 0)
}
