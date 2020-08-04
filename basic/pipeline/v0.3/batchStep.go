package v0_3

import (
	"context"
	"sync"
)

type BatchStep struct {
	Name       string
	process    func(in []interface{}) (out []interface{}, err error)
	fail       func(err error)
	log        func(format string, args ...interface{})
	ctx        context.Context
	inCh       chan interface{}
	outChs     []chan interface{}
	wg         *sync.WaitGroup
	concurrent int
	batchCount int
}

func NewBatchStep(
	name string,
	concurrent int,
	batchCount int,
	process func(in []interface{}) (out []interface{}, err error),
	fail func(err error),
	log func(format string, args ...interface{}),
	ctx context.Context,
	inCh chan interface{},
	outChs []chan interface{},
	wg *sync.WaitGroup,
) *BatchStep {
	return &BatchStep{Name: name, concurrent: concurrent, batchCount: batchCount, process: process, fail: fail, log: log, ctx: ctx, inCh: inCh, outChs: outChs, wg: wg}
}

func (s *BatchStep) Exec() {
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
			var batchItems []interface{}

			defer func() {
				innerWG.Done()

				s.log("Step %s goroutine %d: end exec", s.Name, index)
			}()

			for {
				select {
				case item, ok := <-s.inCh:
					if !ok {
						s.log("Step %s goroutine %d: inCh closed", s.Name, index)

						if len(batchItems) > 0 {
							results, err := s.process(batchItems)
							if err != nil {
								s.fail(err)
								return
							}
							for _, result := range results {
								for _, out := range s.outChs {
									if len(out) == cap(out) {
										s.log("Step %s goroutine %d: len(out)=%d", s.Name, index, len(out))
									}
									out <- result
								}
							}
						}

						return
					}

					batchItems = append(batchItems, item)
					if len(batchItems) == s.batchCount {
						results, err := s.process(batchItems)
						if err != nil {
							s.fail(err)
							return
						}
						//s.log("Step %s goroutine %d: process elapsed %f seconds, len(inCh)=%d", s.Name, index, time.Now().Sub(start).Seconds(), len(s.inCh))

						for _, result := range results {
							for _, out := range s.outChs {
								if len(out) == cap(out) {
									s.log("Step %s goroutine %d: len(out)=%d", s.Name, index, len(out))
								}
								out <- result
							}
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
