package model

import "context"

type PipelineContext struct {
	Ctx     context.Context
	ErrorCh chan string
}
