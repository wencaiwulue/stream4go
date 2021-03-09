package api

import "go/types"

type StringApi interface {
	Parallel() *StringApi
	Distinct() *StringApi
	ToSlice() []string
	Of(elements ...string) *StringApi
	Map(f func(string) string) *StringApi
	MapToObj(f func(string) interface{}) *StringApi
	FlatMap(f func(string) string) *StringApi
	Collect(f func(string) string) *StringApi
	Filter(f func(string) string) *StringApi
	DistinctF(f func(string) string) *StringApi
	Sorted(f func(string) string) *StringApi
	SortedDefault() *StringApi
	Skip(n int64) *StringApi
	Limit(n int64) *StringApi
	Peek(func(string)) *StringApi
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
