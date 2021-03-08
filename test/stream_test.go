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

func TestLong(t *testing.T) {
	impl.LongStream.Range(1, 10).ForEach(func(int642 int64) {
		fmt.Println(int642)
	})
}

func TestReflect(t *testing.T) {
	u := User{
		name: Name{a: "asdf"},
		age:  1,
	}
	impl.ObjectStream.
		Of(u, u).
		MapToValue("name").
		MapToValue("a").
		MapToString().
		ForEach(func(string2 string) {
			fmt.Printf("%v\n", string2)
		})
}

func TestCollection(t *testing.T) {
	u := User{
		name: Name{a: "asdf"},
		age:  1,
	}
	impl.ObjectStream.Of([]User{u}, []User{u}).FlatMap().MapToValues("name.a").MapToString().ForEach(func(string2 string) {
		fmt.Println(string2)
	})
}

func TestMap(t *testing.T) {
	u := User{
		name: Name{a: "asdf"},
		age:  1,
	}
	impl.ObjectStream.Of(map[User]User{u: u}, map[User]User{u: u}).FlatMap().MapToValues("key.name.a").MapToString().ForEach(func(string2 string) {
		fmt.Println(string2)
	})
}

type User struct {
	name Name
	age  int64
}
type Name struct {
	a string
}
