package impl

import (
	"math"
	"stream4go/api"
	"stream4go/util"
	"sync"
)

type longStream struct {
	api.LongApi
	elements   []int64
	isParallel bool
}

var longEmpty = &longStream{}

var LongStream = func() *longStream {
	return &longStream{}
}()

func (s *longStream) Of(element ...int64) *longStream {
	return &longStream{
		elements: element,
	}
}

func (s *longStream) Parallel() *longStream {
	s.isParallel = true
	return s
}

func (s *longStream) Map(mapFunc func(s int64) int64) *longStream {
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

func (s *longStream) Range(open int64, close int64) *longStream {
	if close <= open {
		return longEmpty
	}
	temp := make([]int64, 0, close-open)
	for i := open; i < close; i++ {
		temp = append(temp, i)
	}
	s.elements = temp
	return s
}

// TODO
func (s *longStream) MapToObj(mapFunc func(s string) interface{}) *longStream {
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

func (s *longStream) ToSlice() []int64 {
	return s.elements[0:len(s.elements)]
}

func (s *longStream) Skip(offset int64) *longStream {
	if len(s.elements) < int(offset) {
		return longEmpty
	} else {
		s.elements = s.elements[offset:]
		return s
	}
}

func (s *longStream) Limit(limit int64) *longStream {
	if len(s.elements) > int(limit) {
		s.elements = s.elements[0:limit]
	}
	return s
}

func (s *longStream) Count() int64 {
	return int64(len(s.elements))
}

func (s *longStream) Max() int64 {
	max := int64(math.MinInt64)
	for _, e := range s.elements {
		max = util.MaxLong(max, e)
	}
	return max
}

func (s *longStream) Min() int64 {
	min := int64(math.MaxInt64)
	for _, e := range s.elements {
		min = util.MinLong(min, e)
	}
	return min
}

func (s *longStream) Distinct() *longStream {
	set := make(map[int64]int64)
	result := make([]int64, 0, len(set))
	for _, e := range s.elements {
		if set[e] == 0 {
			result = append(result, e)
		}
		set[e] = e
	}
	s.elements = result[0:]
	return s
}

func (s *longStream) Filter(predicate func(string2 int64) bool) *longStream {
	strings := make([]int64, 0, len(s.elements))
	for _, e := range s.elements {
		if predicate(e) {
			strings = append(strings, e)
		}
	}
	s.elements = strings[0:]
	return s
}

func (s *longStream) ForEach(f func(int642 int64)) {
	for _, element := range s.elements {
		f(element)
	}
}
