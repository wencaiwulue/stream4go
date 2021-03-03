package stream4go

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	s := &Stream{}
	result := s.
		Of("a", "b").
		Map(func(i interface{}) interface{} {
			return i
		}).
		ToList()
	for _, e := range result {
		fmt.Println(e)
	}
}
