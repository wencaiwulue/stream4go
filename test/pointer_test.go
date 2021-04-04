package test

import (
	"fmt"
	"testing"
)

type Interface interface {
	function(a int)
}
type Impl struct {
	b string
}

func (b *Impl) function(a int) {
	fmt.Println(a, b.b)
}

func test(t Interface) {
	if t == nil {
		fmt.Println("nil")
	} else {
		t.function(10)
	}
}

func TestMa(t *testing.T) {
	var a Interface
	var b *Impl

	a = b
	fmt.Println(a == nil)
	fmt.Println(a)
	test(a)
	a.function(10)
	//b.function(1)
}
