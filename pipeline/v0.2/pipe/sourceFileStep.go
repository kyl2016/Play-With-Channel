package pipe

import "github.com/kyl2016/Play-With-Golang/pipeline/v0.2/model"

type SourceFileStep struct {
	files []string
}

func NewSourceFileStep(files []string) *SourceFileStep {
	return &SourceFileStep{
		files: files,
	}
}

func (s *SourceFileStep) Exec() chan interface{} {
	fileCh := make(chan interface{}, model.BUFFERSIZE)

	go func() {
		for _, file := range s.files {
			fileCh <- file
		}

		close(fileCh)
	}()

	return fileCh
}
