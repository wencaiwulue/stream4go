package stream

import (
	"reflect"
	"stream4go/api"
)

type entryStream struct {
	api.Api
	elements   []entry
	isParallel bool
}

var entryEmpty = &entryStream{}

var EntryStream = func() *entryStream {
	return &entryStream{}
}()

type entry struct {
	key   reflect.Value
	value reflect.Value
}
