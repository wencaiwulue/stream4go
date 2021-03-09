package api

import "go/types"

type LongApi interface {
	Parallel() *LongApi
	Distinct() *LongApi
	ToSlice() []int64
	Of(elements ...int64) *LongApi
	Range(open int64, close int64) *LongApi
	Map(f func(int642 int64) int64) *LongApi
	FlatMap(f func(int64) int64) *LongApi
	Collect(f func(int64) int64) *LongApi
	Filter(f func(int64) int64) *LongApi
	DistinctF(f func(int64) int64) *LongApi
	Sorted(f func(int64) int64) *LongApi
	SortedDefault() *LongApi
	Skip(n int64) *LongApi
	Limit(n int64) *LongApi
	Peek(func(int64)) *LongApi
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
