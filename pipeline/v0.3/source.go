package v0_3

import (
	"sync"
)

type Source struct {
	Name    string
	process func(p func(in interface{}))
	log     func(format string, args ...interface{})
	fail    func(err error)
	outChs  []chan interface{}
	wg      *sync.WaitGroup
}

func NewSource(name string,
	process func(p func(in interface{})),
	log func(format string, args ...interface{}),
	fail func(err error),
	outChs []chan interface{},
	wg *sync.WaitGroup) *Source {
	return &Source{
		Name:    name,
		process: process,
		log:     log,
		fail:    fail,
		outChs:  outChs,
		wg:      wg,
	}
}

func (s *Source) Exec() {
	s.log("Source %s: begin exec", s.Name)

	s.wg.Add(1)

	go func() {
		defer func() {

			for _, out := range s.outChs {
				close(out)
			}

			s.wg.Done()

			s.log("Source %s: end exec", s.Name)
		}()

		s.process(func(in interface{}) {
			for _, out := range s.outChs {
				s.log("Source %s: len(outChs)=%d", s.Name, len(s.outChs))

				out <- in
			}
		})
	}()
}
