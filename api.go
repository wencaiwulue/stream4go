package stream4go

type Api interface {
	Of(elements ...interface{}) *Stream
	Map(f func(interface{}) interface{}) *Stream
	FlatMap(f func(interface{}) interface{}) *Stream
	Collect(f func(interface{}) interface{}) *Stream
	filter(f func(interface{}) interface{}) *Stream
	distinct(f func(interface{}) interface{}) *Stream
	sorted(f func(interface{}) interface{}) *Stream
	sortedDefault() *Stream
	skip(n int64) *Stream
	limit(n int64) *Stream
	peek() *Stream
	forEach()
	forEachOrdered()
	toArray() []interface{}
	// Generic type
	toArrayA() []interface{}
	reduce() interface{}
	min() interface{}
	max() interface{}
	count() int64
	anyMatch() int64
	noneMatch() int64
	findFirst() interface{}
	findAny() interface{}
}
