package internal

import (
	"fmt"
	"testing"
)

func TestVersion(t *testing.T) {
	v := NewVersion("v0.15")
	fmt.Println(v.GetValue())
}
