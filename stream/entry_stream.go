package stream

import (
	"github.com/wencaiwulue/stream4go/api"
	"reflect"
)

type entryStream struct {
	api.Api
	elements   []Entry
	isParallel bool
}

var entryEmpty = &entryStream{}

var EntryStream = func() *entryStream {
	return &entryStream{}
}()

type Entry struct {
	Key   reflect.Value
	Value reflect.Value
}
