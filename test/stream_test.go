package test

import (
	"fmt"
	"github.com/wencaiwulue/stream4go/stream"
	"reflect"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	slice := stream.StringStream.
		Of([]string{"first", "b", "c"}...).
		Map(func(s string) string { return s + s }).
		Filter(func(s string) bool { return strings.Contains(s, "cc") }).
		Distinct().
		ToSlice()
	for _, s := range slice {
		fmt.Println(s)
	}
}

func TestLong(t *testing.T) {
	stream.LongStream.
		Range(1, 10).
		ForEach(func(int642 int64) { fmt.Println(int642) })
}

func TestReflect(t *testing.T) {
	user := User{
		name: Name{first: "asdf"},
		age:  1,
	}
	stream.ObjectStream.
		Of(user, user).
		MapToValue("name").
		MapToValue("first").
		MapToString().
		ForEach(func(s string) { fmt.Printf("%v\n", s) })
}

func TestCollection(t *testing.T) {
	user := User{
		name: Name{first: "asdf"},
		age:  1,
	}
	stream.ObjectStream.
		Of([]User{user}, []User{user}).
		FlatMap().
		MapToValues("name.first").
		MapToString().
		ForEach(func(s string) { fmt.Println(s) })
}

func TestMap(t *testing.T) {
	user := User{
		name: Name{first: "naison"},
		age:  1,
	}
	stream.ObjectStream.
		Of(map[User]User{user: user}, map[User]User{user: user}).
		FlatMap().
		//MapToValues("key.name.first").
		MapToValue("key").
		MapToValue("name").
		MapToValue("first").
		MapTo(reflect.TypeOf(""))
}

type User struct {
	name Name
	age  int64
}
type Name struct {
	first string
}
