package framework

import (
	"context"
	"fmt"
)

type SourceStep struct {
	Name             string
	context          context.Context
	AfterAllFinished func()
	Prepare          func(pipeContext interface{}, in interface{}, out chan interface{})
	LogCh            chan string
}

func (ss *SourceStep) Exec(businessContext interface{}, in interface{}) chan interface{} {
	out := make(chan interface{}, BUFFERSIZE)

	writeLog(ss.LogCh, fmt.Sprintf("begin source step: %s", ss.Name))

	go func() {
		defer func() {
			if ss.AfterAllFinished != nil {
				ss.AfterAllFinished()
			}

			close(out)

			writeLog(ss.LogCh, fmt.Sprintf("end source step: %s", ss.Name))
		}()

		ss.Prepare(businessContext, in, out)
	}()

	return out
}
