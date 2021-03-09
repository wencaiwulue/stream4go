package stream

import (
	"stream4go/api"
	"stream4go/util"
	"sync"
)

type stringStream struct {
	api.StringApi
	elements   []string
	isParallel bool
}

var emptyString = &stringStream{}

var StringStream = func() *stringStream {
	return &stringStream{}
}()

func (s *stringStream) Of(element ...string) *stringStream {
	return &stringStream{
		elements: element,
	}
}

func (s *stringStream) Parallel() *stringStream {
	s.isParallel = true
	return s
}

func (s *stringStream) Map(mapFunc func(s string) string) *stringStream {
	if s.isParallel {
		wait := sync.WaitGroup{}
		wait.Add(len(s.elements))
		for i := 0; i < len(s.elements); i++ {
			go func(i int) {
				s.elements[i] = mapFunc(s.elements[i])
				wait.Done()
			}(i)
		}
		wait.Wait()
	} else {
		for i := 0; i < len(s.elements); i++ {
			s.elements[i] = mapFunc(s.elements[i])
		}
	}
	return s
}

// TODO
func (s *stringStream) MapToObj(mapFunc func(s string) interface{}) *stringStream {
	if s.isParallel {
		wait := sync.WaitGroup{}
		wait.Add(len(s.elements))
		for i := 0; i < len(s.elements); i++ {
			go func(i int) {
				//s.elements[i] = mapFunc(s.elements[i])
				wait.Done()
			}(i)
		}
		wait.Wait()
	} else {
		for i := 0; i < len(s.elements); i++ {
			//s.elements[i] = mapFunc(s.elements[i])
		}
	}
	return s
}

func (s *stringStream) ToSlice() []string {
	return s.elements[0:len(s.elements)]
}

func (s *stringStream) Skip(offset int64) *stringStream {
	if len(s.elements) < int(offset) {
		return emptyString
	} else {
		s.elements = s.elements[offset:]
		return s
	}
}

func (s *stringStream) Limit(limit int64) *stringStream {
	if len(s.elements) > int(limit) {
		s.elements = s.elements[0:limit]
	}
	return s
}

func (s *stringStream) Count() int64 {
	return int64(len(s.elements))
}

func (s *stringStream) Max() string {
	max := ""
	for _, e := range s.elements {
		max = util.MaxString(max, e)
	}
	return max
}

func (s *stringStream) Min() string {
	min := ""
	for _, e := range s.elements {
		min = util.MinString(min, e)
	}
	return min
}

func (s *stringStream) Distinct() *stringStream {
	set := make(map[string]string)
	result := make([]string, 0, len(set))
	for _, e := range s.elements {
		if set[e] == "" {
			result = append(result, e)
		}
		set[e] = e
	}
	s.elements = result[0:]
	return s
}

func (s *stringStream) Filter(predicate func(string2 string) bool) *stringStream {
	strings := make([]string, 0, len(s.elements))
	for _, e := range s.elements {
		if predicate(e) {
			strings = append(strings, e)
		}
	}
	s.elements = strings[0:]
	return s
}

func (s *stringStream) ForEach(f func(string2 string)) {
	for _, element := range s.elements {
		f(element)
	}
}
