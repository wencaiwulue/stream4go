package stream

type Api interface {
	Of(elements ...interface{}) *Stream
	Map(f func(interface{}) interface{}) *Stream
	Collect(f func(interface{}) interface{}) *Stream
}

type Stream struct {
	Api
	Element []interface{}
}

var Streams = func() *Stream {
	return &Stream{}
}()

func (s *Stream) Of(elements ...interface{}) *Stream {
	return &Stream{Element: elements}
}

func (s *Stream) Map(f func(interface{}) interface{}) *Stream {
	temp := make([]interface{}, 0, len(s.Element))
	for _, e := range s.Element {
		temp = append(temp, f(e))
	}
	s.Element = temp
	return s
}

func (s *Stream) Collect(f func(interface{}) interface{}) {

}

func (s *Stream) ToList() []interface{} {
	return s.Element
}
