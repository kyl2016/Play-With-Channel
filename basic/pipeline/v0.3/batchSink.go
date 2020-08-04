package v0_3

import (
	"context"
	"sync"
)

type BatchSink struct {
	Name       string
	concurrent int
	batchCount int
	process    func(in []interface{}) (err error)
	fail       func(err error)
	log        func(format string, args ...interface{})
	ctx        context.Context
	inCh       chan interface{}
	wg         *sync.WaitGroup
}

func NewBatchSink(
	name string,
	concurrent int,
	batchCount int,
	process func(in []interface{}) (err error),
	fail func(err error),
	log func(format string, args ...interface{}),
	ctx context.Context,
	inCh chan interface{},
	wg *sync.WaitGroup) *BatchSink {
	return &BatchSink{Name: name, concurrent: concurrent, batchCount: batchCount, process: process, fail: fail, log: log, ctx: ctx, inCh: inCh, wg: wg}
}

func (s *BatchSink) Exec() {
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
			var batchItems []interface{}

			defer func() {
				innerWG.Done()

				s.log("Sink %s goroutine %d: end exec", s.Name, index)
			}()

			for {
				select {
				case item, ok := <-s.inCh:
					if !ok {
						s.log("Sink %s goroutine %d: inCh closed", s.Name, index)

						if len(batchItems) > 0 {
							if err := s.process(batchItems); err != nil {
								s.fail(err)
								return
							}
							batchItems = batchItems[:0]
						}

						return
					}

					batchItems = append(batchItems, item)
					if len(batchItems) == s.batchCount {
						if err := s.process(batchItems); err != nil {
							s.fail(err)
							return
						}

						batchItems = batchItems[:0]
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
