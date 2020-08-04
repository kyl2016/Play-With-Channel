package pipe

func Exec() {
	files := []string{"jsfl", "sdlfjsfd"}
	source := NewSourceFileStep(files)
	out := source.Exec()

	createTask := createExtractTaskStep{}
	out2 := createTask.Exec(out)

}
