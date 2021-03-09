package impl

import (
	"stream4go/api"
	"stream4go/util"
	"sync"
)

type StrStream struct {
	api.StringApi
	elements   []string
	isParallel bool
}

var empty = &StrStream{}

var StringStream = func() *StrStream {
	return &StrStream{}
}()

func (s *StrStream) Of(element ...string) *StrStream {
	return &StrStream{
		elements: element,
	}
}

func (s *StrStream) Parallel() *StrStream {
	s.isParallel = true
	return s
}

func (s *StrStream) Map(mapFunc func(s string) string) *StrStream {
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
func (s *StrStream) MapToObj(mapFunc func(s string) interface{}) *StrStream {
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

func (s *StrStream) ToSlice() []string {
	return s.elements[0:len(s.elements)]
}

func (s *StrStream) Skip(offset int64) *StrStream {
	if len(s.elements) < int(offset) {
		return empty
	} else {
		s.elements = s.elements[offset:]
		return s
	}
}

func (s *StrStream) Limit(limit int64) *StrStream {
	if len(s.elements) > int(limit) {
		s.elements = s.elements[0:limit]
	}
	return s
}

func (s *StrStream) Count() int64 {
	return int64(len(s.elements))
}

func (s *StrStream) Max() string {
	max := ""
	for _, e := range s.elements {
		max = util.MaxString(max, e)
	}
	return max
}

func (s *StrStream) Min() string {
	min := ""
	for _, e := range s.elements {
		min = util.MinString(min, e)
	}
	return min
}

func (s *StrStream) Distinct() *StrStream {
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

func (s *StrStream) Filter(predicate func(string2 string) bool) *StrStream {
	strings := make([]string, 0, len(s.elements))
	for _, e := range s.elements {
		if predicate(e) {
			strings = append(strings, e)
		}
	}
	s.elements = strings[0:]
	return s
}

func (s *StrStream) ForEach(f func(string2 string)) {
	for _, element := range s.elements {
		f(element)
	}
}
