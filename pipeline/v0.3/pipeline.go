package v0_3

import (
	"context"
	"sync"
)

type Pipeline struct {
	Name     string
	State    PipelineState
	stages   []*Stage // stages execute sequently
	failFast bool     // fail when some stages failed
	ctx      context.Context
	cancel   context.CancelFunc
	log      func(format string, args ...interface{})
}

func NewPipeline(name string, failFast bool, log func(format string, args ...interface{})) *Pipeline {
	ctx, cancel := context.WithCancel(context.Background())

	p := Pipeline{
		Name:     name,
		failFast: failFast,
		ctx:      ctx,
		cancel:   cancel,
		log:      log,
	}

	return &p
}

func (p *Pipeline) Run() {
	p.State = Running

	for _, stage := range p.stages {
		stage.Run() // FIXME: ok?
	}

	if p.State == Running {
		p.State = Finished
	}
}

func (p *Pipeline) Stop() {
	p.State = Canceled
	p.cancel()
}

func (p *Pipeline) fail(err error) {
	if p.failFast {
		p.log("pipeline %s: fail with error:%s, will cancel", p.Name, err.Error())
		p.State = Failed
		p.cancel()
	}
}

func (p *Pipeline) NewStage(name string) *Stage {
	s := &Stage{
		Name: name,
		ctx:  p.ctx,
		log:  p.log,
		wg:   &sync.WaitGroup{},
		fail: p.fail,
	}

	p.stages = append(p.stages, s)

	return s
}

type PipelineState string

const (
	Running  PipelineState = "RUNNING"
	Canceled PipelineState = "CANCELED"
	Failed   PipelineState = "FAILED"
	Finished PipelineState = "FINISHED"
)
