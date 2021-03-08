package api

import "go/types"

type StringApi interface {
	Parallel() *Api
	Distinct() *Api
	ToSlice() []string
	Of(elements ...string) *Api
	Map(f func(string) string) *Api
	MapToObj(f func(string) interface{}) *Api
	FlatMap(f func(string) string) *Api
	Collect(f func(string) string) *Api
	Filter(f func(string) string) *Api
	DistinctF(f func(string) string) *Api
	Sorted(f func(string) string) *Api
	SortedDefault() *Api
	Skip(n int64) *Api
	Limit(n int64) *Api
	Peek(func(string)) *Api
	ForEach(func(string))
	ForEachOrdered()
	ToArray() []string
	// Generic type
	ToArrayA() []string
	Reduce() string
	Min() string
	Max() string
	Count() int64
	AnyMatch() int64
	NoneMatch() int64
	FindFirst() string
	FindAny(array types.Type) string
}
