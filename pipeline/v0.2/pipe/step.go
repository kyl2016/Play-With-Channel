package pipe

import (
	"github.com/kyl2016/Play-With-Golang/pipeline/v0.1/pipe/model"
	"sync"
)

type Step struct {
	Name              string
	Process           func(ctx *model.PipelineContext)
	Concurrent        int
	AfterAllCompleted func()
	pipeContext       *model.PipelineContext
	wg                *sync.WaitGroup
}
