package main

func main() {
	print2(nil)
}

func print2(s *myStruct2) {
	println(s.Name())
}

type myStruct2 struct {
	name string
}

func (s *myStruct2) Name() string {
	if s == nil {
		return ""
	}

	return s.name
}
