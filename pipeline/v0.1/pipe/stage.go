package pipe

import (
	"fmt"
	"github.com/kyl2016/Play-With-Golang/pipeline/v0.1/pipe/model"
	"sync"
)

type Stage struct {
	Name     string
	steps    []Step
	Ctx      *model.PipelineContext
	canceled bool
}

func (s *Stage) Exec() {
	fmt.Println("begin exec stage ", s.Name)

	go func() {
		<-s.Ctx.Ctx.Done()
		s.canceled = true
	}()

	wg := &sync.WaitGroup{}

	for _, step := range s.steps {
		if s.canceled {
			break
		}

		step.wg = wg
		step.pipeContext = s.Ctx

		func(step Step) {
			step.Exec()
		}(step)
	}

	wg.Wait()

	fmt.Println("end exec stage ", s.Name)

	return
}

func (s *Stage) AddStep(step Step) {
	s.steps = append(s.steps, step)
}
