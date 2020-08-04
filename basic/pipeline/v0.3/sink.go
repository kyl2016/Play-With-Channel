package v0_3

import (
	"context"
	"sync"
)

type Sink struct {
	Name       string
	concurrent int
	process    func(in interface{}) (err error)
	fail       func(err error)
	log        func(format string, args ...interface{})
	ctx        context.Context
	inCh       chan interface{}
	wg         *sync.WaitGroup
}

func NewSink(
	name string,
	concurrent int,
	process func(in interface{}) (err error),
	fail func(err error),
	log func(format string, args ...interface{}),
	ctx context.Context,
	inCh chan interface{},
	wg *sync.WaitGroup) *Sink {
	return &Sink{Name: name, concurrent: concurrent, process: process, fail: fail, log: log, ctx: ctx, inCh: inCh, wg: wg}
}

func (s *Sink) Exec() {
	s.log("Sink %s: begin %d concurrent exec", s.Name, s.concurrent)

	s.wg.Add(1)

	innerWG := &sync.WaitGroup{}

	go func() {
		innerWG.Wait()

		s.wg.Done()

		s.log("Sink %s: end exec", s.Name)
	}()

	for i := 0; i < s.concurrent; i++ {
		innerWG.Add(1)

		go func(index int) {
			defer func() {
				innerWG.Done()

				s.log("Sink %s goroutine %d: end exec", s.Name, index)
			}()

			for {
				select {
				case item, ok := <-s.inCh:
					if !ok {
						s.log("Sink %s goroutine %d: inCh closed", s.Name, index)
						return
					}

					if err := s.process(item); err != nil {
						s.fail(err)
						return
					}

					//s.log("Sink %s goroutine %d: process elapsed %f seconds, len(inCh)=%d", s.Name, index, time.Now().Sub(start).Seconds(), len(s.inCh))
				case <-s.ctx.Done():
					s.log("Sink %s goroutine %d: canceled", s.Name, index)
					return
				}
			}
		}(i)
	}
}
