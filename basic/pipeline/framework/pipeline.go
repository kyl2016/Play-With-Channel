package framework

import "context"

type Pipeline struct {
	businessContext interface{}     // business logic context
	context         context.Context // context for cancel all goroutines
	SourceStep      SourceStep      // the first step, provide an output channel
	Steps           []Step          // multi steps
	SinkStep        SinkStep        // the last step
}

func NewPipeline(ctx context.Context, context interface{}) *Pipeline {
	return &Pipeline{businessContext: context, context: ctx}
}

func (p *Pipeline) Run(in interface{}) chan interface{} {
	out := p.SourceStep.Exec(p.businessContext, in)

	for _, step := range p.Steps {
		func(s Step) {
			out = s.Exec(p.context, p.businessContext, out)
		}(step)
	}

	return p.SinkStep.Exec(p.businessContext, out)
}
