package api

import "go/types"

type baseApi interface {
	isParallel() bool
	sequential() *baseApi
	parallel() *baseApi
	unordered() *baseApi
	onClose(closeHandler func()) *baseApi
	close()
}

type Api interface {
	baseApi

	Distinct() *Api
	ToSlice() []interface{}
	Of(elements ...interface{}) *Api
	Map(f func(interface{}) interface{}) *Api
	FlatMap(f func(interface{}) interface{}) *Api
	Collect(f func(interface{}) interface{}) *Api
	Filter(f func(interface{}) interface{}) *Api
	DistinctF(f func(interface{}) interface{}) *Api
	Sorted(f func(interface{}) interface{}) *Api
	SortedDefault() *Api
	Skip(n int64) *Api
	Limit(n int64) *Api
	Peek() *Api
	ForEach()
	ForEachOrdered()
	ToArray() []interface{}
	// Generic type
	ToArrayA() []interface{}
	Reduce() interface{}
	Min() interface{}
	Max() interface{}
	Count() int64
	AnyMatch() int64
	NoneMatch() int64
	FindFirst() interface{}
	FindAny(array types.Type) interface{}
}
