package pubsub

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Subscriber interface {
	Notify(interface{}) error
	Close()
}
type writerSubscriber struct {
	in     chan interface{}
	id     int
	Writer io.Writer
}

func (s *writerSubscriber) Notify(msg interface{}) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%#v", rec)
		}
	}()

	select {
	case s.in <- msg:
	case <-time.After(time.Second):
		err = fmt.Errorf("timeout")
	}
	return
}

func (s *writerSubscriber) Close() {
	close(s.in)
}

func NewWriterSubscriber(id int, out io.Writer) Subscriber {
	if out == nil {
		out = os.Stdout
	}

	s := &writerSubscriber{
		id:     id,
		in:     make(chan interface{}),
		Writer: out,
	}

	go func() {
		for msg := range s.in {
			fmt.Fprintf(s.Writer, "(W%d): %v\n", s.id, msg)
		}
	}()
	return s
}
