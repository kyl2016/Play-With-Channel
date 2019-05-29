package framework

type ErrorInfo struct {
	Type string
	Info string
}

type Status struct {
	ErrorInfo
	TaskID int
	Code int
}