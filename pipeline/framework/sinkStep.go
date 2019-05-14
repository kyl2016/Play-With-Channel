package framework

import (
	"context"
	"fmt"
)

type SinkStep struct {
	Name    string
	context context.Context
	Process func(pipeContext interface{}, in interface{}) interface{}
	LogCh   chan string
}

func (s *SinkStep) Exec(businessContext interface{}, in chan interface{}) chan interface{} {
	out := make(chan interface{}, BUFFERSIZE)

	writeLog(s.LogCh, fmt.Sprintf("begin sink step: %s", s.Name))

	go func() {
		defer func() {
			close(out)

			writeLog(s.LogCh, fmt.Sprintf("end sink step: %s", s.Name))
		}()

		for item := range in {
			r := s.Process(businessContext, item)

			// check out is full
			if len(out) == BUFFERSIZE {
				writeLog(s.LogCh, fmt.Sprintf("output channel is full of %s", s.Name))
			}

			out <- r
		}
	}()

	return out
}
