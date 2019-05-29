package pipe

import "github.com/kyl2016/Play-With-Golang/pipeline/v0.1/pipe/model"

type Pipeline struct {
	stages   []Stage
	ctx      *model.PipelineContext
	canceled bool
}

func NewPipeline(ctx *model.PipelineContext) *Pipeline {
	return &Pipeline{ctx: ctx}
}

func (p *Pipeline) Exec() {
	go func() {
		<-p.ctx.Ctx.Done()
		p.canceled = true
	}()

	for _, s := range p.stages {
		if !p.canceled {
			s.Exec()
		}
	}
}

func (p *Pipeline) AddStage(s ...Stage) {
	p.stages = append(p.stages, s...)
}
