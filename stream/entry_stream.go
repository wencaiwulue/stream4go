package stream

import (
	"github.com/wencaiwulue/stream4go/api"
	"reflect"
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
