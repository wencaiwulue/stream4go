package api

import "go/types"

type LongApi interface {
	Parallel() *Api
	Distinct() *Api
	ToSlice() []int64
	Of(elements ...int64) *Api
	Range(open int64, close int64) *Api
	Map(f func(int642 int64) int64) *Api
	FlatMap(f func(int64) int64) *Api
	Collect(f func(int64) int64) *Api
	Filter(f func(int64) int64) *Api
	DistinctF(f func(int64) int64) *Api
	Sorted(f func(int64) int64) *Api
	SortedDefault() *Api
	Skip(n int64) *Api
	Limit(n int64) *Api
	Peek(func(int64)) *Api
	ForEach(f func(int642 int64))
	ForEachOrdered()
	ToArray() []int64
	// Generic type
	ToArrayA() []int64
	Reduce() int64
	Min() int64
	Max() int64
	Count() int64
	AnyMatch() int64
	NoneMatch() int64
	FindFirst() int64
	FindAny(array types.Type) int64
}
