package v0_3

import (
	"context"
	"sync"
	"time"
)

type Step struct {
	Name       string
	process    func(in interface{}) (out interface{}, err error)
	fail       func(err error)
	log        func(format string, args ...interface{})
	ctx        context.Context
	inCh       chan interface{}
	outChs     []chan interface{}
	wg         *sync.WaitGroup
	concurrent int
	batchCount int
}

func NewStep(
	name string,
	concurrent int,
	process func(in interface{}) (out interface{}, err error),
	fail func(err error),
	log func(format string, args ...interface{}),
	ctx context.Context,
	inCh chan interface{},
	outChs []chan interface{},
	wg *sync.WaitGroup,
) *Step {
	return &Step{Name: name, concurrent: concurrent, process: process, fail: fail, log: log, ctx: ctx, inCh: inCh, outChs: outChs, wg: wg}
}

func (s *Step) Exec() {
	s.log("Step %s: begin %d concurrent exec", s.Name, s.concurrent)

	s.wg.Add(1)

	innerWG := &sync.WaitGroup{}

	go func() {
		innerWG.Wait()

		for _, out := range s.outChs {
			close(out)
			s.log("Step %s: close outCh %p", s.Name, out)
		}

		s.wg.Done()

		s.log("Step %s: end exec", s.Name)
	}()

	for i := 0; i < s.concurrent; i++ {
		innerWG.Add(1)

		go func(index int) {
			defer func() {
				innerWG.Done()

				s.log("Step %s goroutine %d: end exec", s.Name, index)
			}()

			for {
				select {
				case item, ok := <-s.inCh:
					if !ok {
						s.log("Step %s goroutine %d: inCh closed", s.Name, index)
						return
					}

					start := time.Now()
					r, err := s.process(item)
					if err != nil {
						s.fail(err)
						return
					}
					s.log("Step %s goroutine %d: process elapsed %f seconds, len(inCh)=%d", s.Name, index, time.Now().Sub(start).Seconds(), len(s.inCh))

					if r != nil {
						for _, out := range s.outChs {
							if len(out) == cap(out) {
								s.log("Step %s goroutine %d: len(out)=%d", s.Name, index, len(out))
							}
							out <- r
						}
					}
				case <-s.ctx.Done():
					s.log("Step %s goroutine %d: canceled", s.Name, index)
					return
				}
			}
		}(i)
	}
}
