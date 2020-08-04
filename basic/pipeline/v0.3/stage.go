package v0_3

import (
	"context"
	"fmt"
	"sync"
)

type Stage struct {
	Name   string
	source Executor
	steps  []Executor
	sink   Executor
	ctx    context.Context
	log    func(format string, args ...interface{})
	wg     *sync.WaitGroup
	fail   func(err error)
}

func (s *Stage) Run() {
	if s.source == nil {
		s.fail(fmt.Errorf("Stage %s: didn't set Source", s.Name))
		return
	}

	s.log("Stage %s: begin run", s.Name)

	s.source.Exec()

	for _, step := range s.steps {
		func(s Executor) {
			s.Exec()
		}(step)
	}

	if s.sink != nil {
		s.sink.Exec()
	}

	s.wg.Wait()

	s.log("Stage %s: end run", s.Name)
}

func (s *Stage) SetSource(name string, process func(p func(in interface{})), outChs ...chan interface{}) {
	s.source = NewSource(name, process, s.log, s.fail, outChs, s.wg)
}

func (s *Stage) AddStep(name string, concurrent int, process func(in interface{}) (out interface{}, err error), inCh chan interface{}, outChs ...chan interface{}) {
	s.steps = append(s.steps, NewStep(name, concurrent, process, s.fail, s.log, s.ctx, inCh, outChs, s.wg))
}

func (s *Stage) AddBatchStep(name string, concurrent, batchCount int, process func(in []interface{}) (out []interface{}, err error), inCh chan interface{}, outChs ...chan interface{}) {
	s.steps = append(s.steps, NewBatchStep(name, concurrent, batchCount, process, s.fail, s.log, s.ctx, inCh, outChs, s.wg))
}

func (s *Stage) SetSink(name string, concurrent int, process func(in interface{}) (err error), inCh chan interface{}) {
	s.sink = NewSink(name, concurrent, process, s.fail, s.log, s.ctx, inCh, s.wg)
}

func (s *Stage) SetBatchSink(name string, concurrent, batchCount int, process func(in []interface{}) (err error), inCh chan interface{}) {
	s.sink = NewBatchSink(name, concurrent, batchCount, process, s.fail, s.log, s.ctx, inCh, s.wg)
}
