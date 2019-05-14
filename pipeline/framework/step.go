package framework

import (
	"context"
	"fmt"
)

const (
	BUFFERSIZE = 100
)

type Step struct {
	Name             string
	AfterAllFinished func()
	Process          func(pipeContext interface{}, in interface{}) interface{}
	LogCh            chan string

	//Batch            bool // FIXME: use batch to call process
}

func (s *Step) Exec(context context.Context, businessContext interface{}, in chan interface{}) chan interface{} {
	out := make(chan interface{}, BUFFERSIZE)

	writeLog(s.LogCh, fmt.Sprintf("begin step: %s", s.Name))

	go func() {
		defer func() {
			if s.AfterAllFinished != nil {
				s.AfterAllFinished()
			}

			close(out)
		}()

		for {
			select {
			case item, ok := <-in:
				if !ok {
					writeLog(s.LogCh, fmt.Sprintf("end step: %s", s.Name))
					return
				}

				r := s.Process(businessContext, item)

				// FIXME: check is use batch

				if r != nil {
					// check out is full
					if len(out) == BUFFERSIZE {
						writeLog(s.LogCh, fmt.Sprintf("output channel is full of %s", s.Name))
					}

					out <- r
				}

			case <-context.Done():
				writeLog(s.LogCh, fmt.Sprintf("end step: %s because cancel", s.Name))
				return
			}
		}
	}()

	return out
}
