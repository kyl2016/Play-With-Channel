package main

import (
	"context"
	"fmt"
	"github.com/kyl2016/Play-With-Golang/pipeline/pipe"
	"github.com/kyl2016/Play-With-Golang/pipeline/pipe/model"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	errCh := make(chan string, 100)
	pipeCtx := &model.PipelineContext{
		Ctx:     ctx,
		ErrorCh: errCh,
	}

	stage := pipe.Stage{Name: "stage 1", Ctx: pipeCtx}

	stage.AddStep(pipe.Step{
		Name: "1",
		Process: func(ctx *model.PipelineContext) {
			fmt.Println("process...")
			ctx.ErrorCh <- "1 finished"

			select {
			case <-ctx.Ctx.Done():
				fmt.Println("canceled")
			case <-time.After(time.Second * 3):
				fmt.Println("timeout")
			}

		},
		Concurrent: 1,
		AfterAllCompleted: func() {
			fmt.Println("1 all completed")
		},
	})

	stage.AddStep(pipe.Step{
		Name: "2",
		Process: func(ctx *model.PipelineContext) {
			fmt.Println("process...")
			ctx.ErrorCh <- "2 failed"
		},
		Concurrent: 1,
		AfterAllCompleted: func() {
			fmt.Println("2 all completed")
		},
	})

	stage2 := pipe.Stage{Name: "stage 2", Ctx: pipeCtx}
	stage2.AddStep(pipe.Step{
		Name: "3",
		Process: func(ctx *model.PipelineContext) {
			fmt.Println("process 3...")
			ctx.ErrorCh <- "3 finished"
		},
		Concurrent: 1,
		AfterAllCompleted: func() {
			fmt.Println("all completed")
		},
	})

	go func() {
		for err := range errCh {
			fmt.Println("error info:", err)
			if err == "2 failed" {
				cancel()
			}
		}
	}()

	pipeline := pipe.NewPipeline(pipeCtx)
	pipeline.AddStage(stage, stage2)

	pipeline.Exec()

	time.Sleep(time.Second * 30)
}
