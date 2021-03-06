package test

import (
	"fmt"
	"stream4go/impl"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	ss := impl.StringStream.
		Of([]string{"a", "b", "c"}...).
		Map(func(s string) string { return s + s }).
		Filter(func(s string) bool { return strings.Contains(s, "cc") }).
		Distinct().
		ToSlice()
	for _, e := range ss {
		fmt.Println(e)
	}
}
