package impl

import (
	"stream4go/api"
	"sync"
)

type Stream struct {
	api.StringApi
	element    []string
	isParallel bool
}

var StringStream = func() *Stream {
	return &Stream{}
}()

func (s *Stream) Of(element ...string) *Stream {
	return &Stream{
		element: element,
	}
}

func (s *Stream) parallel() *Stream {
	s.isParallel = true
	return s
}

func (s *Stream) Map(mapFunc func(s string) string) *Stream {
	if s.isParallel {
		wait := sync.WaitGroup{}
		wait.Add(len(s.element))
		for i := 0; i < len(s.element); i++ {
			fi := i
			go func(i int) {
				s.element[i] = mapFunc(s.element[i])
				wait.Done()
			}(fi)
		}
		wait.Wait()
		return s
	} else {
		for i := 0; i < len(s.element); i++ {
			s.element[i] = mapFunc(s.element[i])
		}
	}
	return s
}

func (s *Stream) ToSlice() []string {
	return s.element[0:len(s.element)]
}

func (s *Stream) Distinct() *Stream {
	distinct := make(map[string]string)
	for _, e := range s.element {
		distinct[e] = e
	}
	result := make([]string, 0, len(distinct))
	for k, _ := range distinct {
		result = append(result, k)
	}
	s.element = result
	return s
}

func (s *Stream) Filter(predicate func(string2 string) bool) *Stream {
	strings := make([]string, 0, len(s.element))
	for _, e := range s.element {
		if predicate(e) {
			strings = append(strings, e)
		}
	}
	s.element = strings
	return s
}
