package stream

import (
	"reflect"
	"stream4go/api"
)

type valueStream struct {
	api.Api
	elements   []reflect.Value
	isParallel bool
}

var valueEmpty = &valueStream{}

var ValueStream = func() *valueStream {
	return &valueStream{}
}()

func (s *valueStream) MapToValue(fieldName string) *valueStream {
	values := make([]reflect.Value, 0, len(s.elements))
	for i := range s.elements {
		values = append(values, s.elements[i].FieldByName(fieldName))
	}
	return &valueStream{
		elements: values,
	}
}

func (s *valueStream) MapToString() *stringStream {
	strings := make([]string, 0, len(s.elements))
	for _, element := range s.elements {
		strings = append(strings, element.String())
	}
	return &stringStream{
		elements:   strings,
		isParallel: false,
	}
}

func (s *valueStream) MapToLong() *longStream {
	int64s := make([]int64, 0, len(s.elements))
	for _, element := range s.elements {
		int64s = append(int64s, element.Int())
	}
	return &longStream{
		elements:   int64s,
		isParallel: false,
	}
}
