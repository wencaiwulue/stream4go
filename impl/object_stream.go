package impl

import (
	"fmt"
	"reflect"
	"stream4go/api"
	"strings"
)

type objectStream struct {
	api.Api
	elements   []interface{}
	isParallel bool
}

var objectEmpty = &objectStream{}

var ObjectStream = func() *objectStream {
	return &objectStream{}
}()

func (s *objectStream) Of(element ...interface{}) *objectStream {
	return &objectStream{
		elements: element,
	}
}

func (s *objectStream) FlatMap() *objectStream {
	result := make([]interface{}, 0)
	for _, i := range s.elements {
		k := reflect.ValueOf(i)
		switch k.Kind() {
		case reflect.Slice:
			for i := 0; i < k.Len(); i++ {
				result = append(result, k.Index(i).Interface())
			}
		case reflect.Map:
			iter := k.MapRange()
			for iter.Next() {
				key := iter.Key()
				value := iter.Value()
				result = append(result, entry{
					key:   key,
					value: value,
				})
			}
		default:
			result = append(result, i)
		}
	}
	return &objectStream{
		elements: result,
	}
}

func (s *objectStream) MapToValue(fieldName string) *valueStream {
	values := make([]reflect.Value, 0, len(s.elements))
	for _, i := range s.elements {
		t := reflect.ValueOf(i)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		if t.Kind() != reflect.Struct {
			fmt.Printf("Not struct, %v\n", t)
			return valEmpty
		}
		values = append(values, t.FieldByName(fieldName))
	}
	return &valueStream{
		elements: values,
	}
}

func (s *objectStream) MapToValues(fieldName string) *valueStream {
	split := strings.Split(fieldName, ".")
	l := len(split)
	if l == 0 {
		return valEmpty
	} else if l == 1 {
		return s.MapToValue(split[0])
	} else {
		r := s.MapToValue(split[0])
		for i := 1; i < l; i++ {
			r = r.MapToValue(split[i])
		}
		return r
	}
}
