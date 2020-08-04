package pipe

type createExtractTaskStep struct {
}

func (s *createExtractTaskStep) Exec(in chan interface{}) chan *task {}) {

tasks := make(chan *taskModels.Task, BUFFERSIZE)


return tasks
}
