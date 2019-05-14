package pipe

import (
	"fmt"
	"github.com/kyl2016/Play-With-Golang/pipeline/pipe/model"
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

func (s *Step) Exec() {
	fmt.Println("begin exec step ", s.Name)
	s.wg.Add(1)

	if s.Concurrent == 1 {
		go func() {
			s.Process(s.pipeContext)
			s.wg.Done()
		}()
	} else {
		s.ConcurrentExec()
	}
}

func (s *Step) ConcurrentExec() {
	innerWG := &sync.WaitGroup{}
	innerWG.Add(s.Concurrent)

	//util.Log(util.DebugLevel, "Concurrent", "", fmt.Sprintf("begin workers %s with %d goroutines", name, concurrency))

	for i := 0; i < s.Concurrent; i++ {
		go func(index int) {
			//util.Log(util.DebugLevel, "Concurrent", "", fmt.Sprintf("start worker %s %d", name, index))

			s.Process(s.pipeContext)
			innerWG.Done()

			//util.Log(util.DebugLevel, "Concurrent", "", fmt.Sprintf("start worker %s %d", name, index))
		}(i)
	}

	go func(name string) {
		innerWG.Wait()

		if s.AfterAllCompleted != nil {
			s.AfterAllCompleted()
		}

		fmt.Println("end exec step ", s.Name, name)

		s.wg.Done()

		//util.Log(util.DebugLevel, "Concurrent", "", fmt.Sprintf("finish worker %s with %d goroutines", name, concurrency))
	}(s.Name)
}
